1. Strategy Pattern : 

Strategy pattern(a.k.a Policy Pattern) is used when we have multiple algorithm for a specific task and client decides the actual implementation to be used at runtime.
We define multiple algorithms and let client application pass the algorithm to be used as a parameter. 

Like in ShoppingCart we pass `PaymentStrategy` as paramater of `pay()` method

2. 