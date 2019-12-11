package StrategyPattern;

public class ShoppingCart {
    private int sum;

    public ShoppingCart(){
        this.sum=100;
    }

    public void pay(PaymentStrategy paymentStrategy){
        paymentStrategy.pay(this.sum);
    }
}
