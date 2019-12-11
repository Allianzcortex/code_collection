package StrategyPattern;

public class TestMainFunction {
    public static void main(String[] args){
        ShoppingCart shoppingCart = new ShoppingCart();
        CreditCardStrategy creditCardStrategy = new CreditCardStrategy("1","2","3","4");
        shoppingCart.pay(creditCardStrategy);

        CashStrategy cashStrategy = new CashStrategy();
        shoppingCart.pay(cashStrategy);

        PaypalStrategy paypalStrategy = new PaypalStrategy("1","2");
        shoppingCart.pay(paypalStrategy);
    }

}
