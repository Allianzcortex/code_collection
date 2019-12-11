package StrategyPattern;

public class CreditCardStrategy implements PaymentStrategy {
    private String name;
    private String cardNumber;
    private String CCV;
    private String dateOfExpiry;

    public CreditCardStrategy(String name, String cardNumber, String CCV, String dateOfExpiry) {
        this.name = name;
        this.cardNumber = cardNumber;
        this.CCV = CCV;
        this.dateOfExpiry = dateOfExpiry;
    }

    public boolean isCreditValid() {
        // In real world, there should be a call into outside service
        return true;
    }

    @Override
    public void pay(int amount) {
        if(isCreditValid())
        System.out.println("Pay by Credit Card " + amount);
    }
}
