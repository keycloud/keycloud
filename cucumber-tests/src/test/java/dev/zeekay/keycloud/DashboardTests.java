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
import org.openqa.selenium.chrome.ChromeOptions;

import static org.junit.Assert.*;

public class DashboardTests {
    private ChromeDriver driver;

    @Before("@WithoutPlugin")
    public void setupChrome(){
        WebDriverManager.chromedriver().setup();
        ChromeOptions chromeOptions = new ChromeOptions();
        chromeOptions.addArguments("--headless", "--no-sandbox", "--disable-dev-shm-usage");
        driver = new ChromeDriver(chromeOptions);
    }

    @Given("^I am on the landing page$")
    public void openLandingPage(){
        driver.get("http://localhost:8080/dashboard/index.html");
    }
    @When("^I type in \"([^\"]*)\" as my username and click register$")
    public void insertUsernameAndRegister(String username) throws Exception{
        Thread.sleep(1000);
        WebElement usernameIn = driver.findElementById("inputUser");
        usernameIn.sendKeys(username);
        driver.findElementById("registerBtn").click();
    }
    @Then("^I will be on the settings page of a new created Account$")
    public void checkSettingsPage() throws Throwable{
        Thread.sleep(1000);
        assertEquals(driver.getCurrentUrl(), "localhost:8080/dashboard/main.html#settings");
    }

    @Given("^I am on my home page in the keycloud dashboard$")
    public void openHomePage() throws Exception{
    }

    @When("^I press the add button$")
    public void pressAddPassword() throws Exception{

    }

    @And("^I fill out the popup$")
    public void fillOutPopup() throws Exception{
    }

    @Then("^I will see a new password added to the list$")
    public void checkPasswordAddedToList() throws Exception{
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
