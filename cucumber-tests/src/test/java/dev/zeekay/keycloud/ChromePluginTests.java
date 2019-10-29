package dev.zeekay.keycloud;
import io.cucumber.junit.CucumberOptions;
import io.cucumber.java.PendingException;
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
import org.openqa.selenium.remote.DesiredCapabilities;

import static org.junit.Assert.*;

public class ChromePluginTests {

    private ChromeDriver driver;
    private final String pluginId = "";


    @Before("@WithPlugin")
    public void setupChromeWithPlugin(){
        WebDriverManager.chromedriver().setup();
        ChromeOptions options = new ChromeOptions();
        options.addArguments("load-extension=/keycloud/plugins/KeyCloudChromePlugin");
        driver = new ChromeDriver(options);

    }

    @When("^I click on the autofill - icon$")
    public void clickAutofill() throws Exception {
        throw new PendingException();
    }

    @When("^I successfully log in$")
    public void login() throws Exception {
        throw new PendingException();
    }

    @Then("^the login for \"([^\"]*)\" should be filled$")
    public void checkFilledField(String page) throws Exception {
        throw new PendingException();
    }

    @Given("^I am on \"([^\"]*)\"$")
    public void gotoPage(String page) throws Exception {
        driver.get(page);
    }
    @After()
    public void closeBrowser() {
        if(driver != null)
            driver.quit();
    }
}
