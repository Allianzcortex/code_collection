package StrategyPattern;

public class CashStrategy implements PaymentStrategy {
    @Override
    public void pay(int amount) {
        System.out.println("Pay by Cash "+amount);
    }
}
