package dev.zeekay.keycloud;
import io.cucumber.junit.CucumberOptions;
import io.cucumber.java.After;
import io.cucumber.java.Before;
import io.cucumber.java.en.And;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;
import io.cucumber.junit.Cucumber;
import org.junit.runner.RunWith;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import io.github.bonigarcia.wdm.WebDriverManager;
import org.openqa.selenium.chrome.ChromeDriverService;
import org.openqa.selenium.chrome.ChromeOptions;

import static org.junit.Assert.*;

public class DashboardTests {
    private ChromeDriver driver;
    private String baseUrl;
    private int numberOfEntries;

    @Before("@WithoutPlugin")
    public void setupChrome(){
        WebDriverManager.chromedriver().setup();
        ChromeOptions chrome_options = new ChromeOptions();
		chrome_options.addArguments("--headless", "--no-sandbox", "--disable-dev-shm-usage");
		driver = new ChromeDriver(chrome_options);
        baseUrl = "http://localhost:8000/";
    }


    // UC_register
    @Given("^I am on the landing page$")
    public void openLandingPage(){
        driver.get(baseUrl +"index.html");
    }

    @When("^I type in \"([^\"]*)\" as my username and click register$")
    public void insertUsernameAndRegister(String username) {
        WebElement usernameIn = driver.findElementById("inputUser");
        usernameIn.sendKeys(username);
        driver.findElementById("registerBtn").click();
    }
    @Then("^I will be on the settings page of a new created Account$")
    public void checkSettingsPage() throws Throwable{
        Thread.sleep(2000);
        assertEquals(baseUrl + "main.html?#settings", driver.getCurrentUrl());
    }


    // UC_AddPassword
    @Given("^I am on my home page in the keycloud dashboard$")
    public void openHomePage() throws Exception{
        driver.get(baseUrl + "main.html#home");
    }

    @When("^I press the add button$")
    public void pressAddPassword() throws Exception{
        numberOfEntries = driver.findElementsByClassName("entry").size();
        driver.findElementById("addEntryBtn").click();
    }

    @And("^I fill out the popup$")
    public void fillOutPopup() throws Exception{
        Thread.sleep(1000);
        driver.findElementById("saveEntryBtn").click();
    }

    @Then("^I will see a new password added to the list$")
    public void checkPasswordAddedToList() throws Exception{
        Thread.sleep(3000);
        assertEquals(numberOfEntries+1, driver.findElementsByClassName("entry").size());
    }



    @When("^I press the remove button for the \"([^\"]*)\" password$")
    public void removePasswordEntry(String password) throws Exception{
    }

    @Then("^The password \"([^\"]*)\" entry is removed from the list$")
    public void checkPasswordRemoved(String password) throws Exception{

    }



    @After()
    public void closeBrowser() {
        if(driver != null)
            driver.quit();
    }

    @When("^I copy the password for \"([^\"]*)\" to clipboard$")
    public void copyPasswordToClipboard(String url) {

    }

    @Then("^I have the password for \"([^\"]*)\" in my clipboard$")
    public void checkPasswordInClipboard(String url) {
    }
}
