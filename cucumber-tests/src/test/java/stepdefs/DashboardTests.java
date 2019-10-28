package stepdefs;
import cucumber.api.CucumberOptions;
import cucumber.api.PendingException;
import cucumber.api.java.After;
import cucumber.api.java.Before;
import cucumber.api.java.en.And;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import cucumber.api.junit.Cucumber;
import org.junit.runner.RunWith;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import io.github.bonigarcia.wdm.WebDriverManager;
import static org.junit.Assert.*;

@RunWith(Cucumber.class)
@CucumberOptions(features = "classpath:Feature")
public class DashboardTests {
    private ChromeDriver driver;

    @Before("@WithoutPlugin")
    public void setupChrome(){
        WebDriverManager.chromedriver().setup();
        driver = new ChromeDriver();
    }

    @Given("I am on the landing page")
    public void openLandingPage(){
        driver.get("http://localhost:8080/index.html");
    }
    @When("I type in \"([^\"]*)\" as my username and click register")
    public void insertUsernameAndRegister(String username){
        WebElement usernameIn = driver.findElementById("username_login");
        usernameIn.sendKeys(username);
        driver.findElementById("register").click();
    }
    @Then("I will be on the settings page of a new created Account")
    public void checkSettingsPage() throws Throwable{
        Thread.sleep(1000);
        assertEquals(driver.getCurrentUrl(), "localhost:8080/main.html#settings");
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
