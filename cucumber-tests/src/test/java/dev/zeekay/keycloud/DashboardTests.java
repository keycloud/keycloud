package dev.zeekay.keycloud;

import com.github.javafaker.Faker;
import com.github.javafaker.service.FakeValuesService;
import io.cucumber.junit.CucumberOptions;
import io.cucumber.java.After;
import io.cucumber.java.Before;
import io.cucumber.java.en.And;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;
import io.cucumber.junit.Cucumber;
import org.junit.runner.RunWith;
import org.openqa.selenium.NoAlertPresentException;
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
    private Faker faker;

    @Before("@WithoutPlugin")
    public void setupChrome() {
        WebDriverManager.chromedriver().setup();
        ChromeOptions chrome_options = new ChromeOptions();
        chrome_options.addArguments("--no-sandbox", "--disable-dev-shm-usage"); // TODO add --headless

        driver = new ChromeDriver(chrome_options);
        baseUrl = "http://localhost:8080/dashboard/";
        faker = new Faker();
    }


    // UC_register
    @Given("^I am on the landing page$")
    public void openLandingPage() {
        driver.get(baseUrl + "login.html");
    }

    @When("^I type in \"([^\"]*)\" as my username and click register$")
    public void insertUsernameAndRegister(String username) throws Exception {
        Thread.sleep(2000);
        WebElement usernameIn = driver.findElementById("inputUser");
        usernameIn.sendKeys(username);
        driver.findElementById("registerBtn").click();
    }

    @Then("^I will be on the overview page of a new created Account$")
    public void checkSettingsPage() throws Throwable {
        Thread.sleep(2000);
        assertEquals(baseUrl + "main.html#overview", driver.getCurrentUrl());
    }


    // UC_AddPassword
    @Given("^I am on my home page in the keycloud dashboard$")
    public void openHomePage() throws Exception {
        driver.get(baseUrl + "main.html#overview");
    }

    @When("^I press the add button$")
    public void pressAddPassword() throws Exception {
        Thread.sleep(2000);
        numberOfEntries = driver.findElementsByClassName("entry").size();
        driver.findElementById("addEntryBtn").click();
    }

    @And("^I fill out the popup$")
    public void fillOutPopup() throws Exception {
        Thread.sleep(1000);
        driver.findElementById("usernameInput").sendKeys(faker.name().username());
        driver.findElementById("urlInput").sendKeys(faker.lorem().word());
        driver.findElementById("genBtn").click();
        Thread.sleep(1500);
        driver.findElementById("saveEntryBtn").click();
    }

    @Then("^I will see a new password added to the list$")
    public void checkPasswordAddedToList() throws Exception {
        Thread.sleep(3000);
        assertEquals(numberOfEntries + 1, driver.findElementsByClassName("entry").size());
    }


    @When("^I press the remove button for the \"([^\"]*)\" password$")
    public void removePasswordEntry(String password) throws Exception {
        pressAddPassword();
        fillOutPopup();
        Thread.sleep(500);
        numberOfEntries = driver.findElementsByClassName("entry").size();
        driver.findElementById("rm3").click();
    }

    @Then("^The password \"([^\"]*)\" entry is removed from the list$")
    public void checkPasswordRemoved(String password) throws Exception {
        Thread.sleep(2000);
        assertEquals(numberOfEntries - 1, driver.findElementsByClassName("entry").size());
    }


    @When("^I copy the password for \"([^\"]*)\" to clipboard$")
    public void copyPasswordToClipboard(String url) throws Exception {
        pressAddPassword();
        fillOutPopup();
        Thread.sleep(500);
        driver.findElementById("cp3").click();
    }

    @Then("^I have the password for \"([^\"]*)\" in my clipboard$")
    public void checkPasswordInClipboard(String url) throws Exception{
        Thread.sleep(2000);
        assertTrue(isAlertPresent());
    }


    @After()
    public void closeBrowser() {
        if (driver != null)
            driver.quit();
    }

    private boolean isAlertPresent() {
        try {
            driver.switchTo().alert();
            return true;
        } catch (NoAlertPresentException Ex) {
            return false;
        }

    }
}
