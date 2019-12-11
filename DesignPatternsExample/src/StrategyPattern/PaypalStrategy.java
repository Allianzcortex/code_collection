package StrategyPattern;

public class PaypalStrategy implements PaymentStrategy {
    private String email;
    private String passWord;

    public PaypalStrategy(String email, String passWord) {
        this.email = email;
        this.passWord = passWord;
    }

    public boolean isPaypalAccountValid() {
        // should make a call to paypal
        return true;
    }

    @Override
    public void pay(int amount) {
        System.out.println("Pay by paypal " + amount);
    }
}
