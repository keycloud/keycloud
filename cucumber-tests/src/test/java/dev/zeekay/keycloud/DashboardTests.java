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
    private static String username = "";
    private static String masterPassword = "";

    @Before("@WithoutPlugin")
    public void setupChrome() {
        WebDriverManager.chromedriver().setup();
        ChromeOptions chrome_options = new ChromeOptions();
        chrome_options.addArguments("--no-sandbox", "--disable-dev-shm-usage", "--headless");

        driver = new ChromeDriver(chrome_options);
        baseUrl = "http://localhost:8080/dashboard/";
        faker = new Faker();
    }
    @Given("^I am logged in$")
    public void login() throws Exception{
        driver.get(baseUrl + "login.html");
        driver.findElementById("inputUsername").sendKeys(username);
        driver.findElementById("inputPassword").sendKeys(masterPassword);
        driver.findElementById("loginBtn").click();
        Thread.sleep(1500);
    }

    // UC_register
    @Given("^I am on the landing page$")
    public void openLandingPage()throws Exception {
        driver.get(baseUrl + "login.html");
    }

    @When("^I type in \"([^\"]*)\" as my username and click register$")
    public void insertUsernameAndRegister(String username) throws Exception {
        DashboardTests.username = username;
        WebElement usernameIn = driver.findElementById("inputUser");
        usernameIn.sendKeys(username);
        driver.findElementById("registerBtn").click();
    }

    @Then("^I will be on the overview page of a new created Account$")
    public void checkOverviewPage() throws Throwable {
        Thread.sleep(1000);
        assertEquals(baseUrl + "main.html#overview", driver.getCurrentUrl());
        //Clicking on the element also tests whether the side was rendered correctly and no - Unauthorized - appears
        driver.findElementById("settingsTab").click();
        masterPassword = driver.findElementById("master-pw").getAttribute("value");
    }


    // UC_AddPassword
    @And("^I am on my home page in the keycloud dashboard$")
    public void openHomePage() throws Exception {
        driver.get(baseUrl + "main.html#overview");
    }

    @When("^I press the add button$")
    public void pressAddPassword() throws Exception {
        numberOfEntries = driver.findElementsByClassName("entry").size();
        driver.findElementById("addEntryBtn").click();
    }

    @And("^I fill out the popup$")
    public void fillOutPopup() throws Exception {
        Thread.sleep(1500);
        driver.findElementById("usernameInput").sendKeys(faker.name().username());
        driver.findElementById("urlInput").sendKeys(faker.lorem().word());
        driver.findElementById("genBtn").click();
        Thread.sleep(500);
        driver.findElementById("saveEntryBtn").click();
    }

    @Then("^I will see a new password added to the list$")
    public void checkPasswordAddedToList() throws Exception {
        Thread.sleep(1500);
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
        Thread.sleep(1500);
        assertEquals(numberOfEntries - 1, driver.findElementsByClassName("entry").size());
    }


    @When("^I copy the password for \"([^\"]*)\" to clipboard$")
    public void copyPasswordToClipboard(String url) throws Exception {
        Thread.sleep(500);
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
