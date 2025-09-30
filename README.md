What is the Factory Pattern?

Creational design pattern → focuses on object creation.

Instead of creating objects directly using &Type{}, you delegate creation to a function or type (factory).

Helps when:

You don’t want to expose complex initialization logic.

You want to decide which implementation to return at runtime.

You want to follow OCP (Open/Closed Principle) — add new products without changing client code.


Simple Factory 

What it is: A function that creates objects based on input parameters.

Purpose: Encapsulate object creation in one place so the client doesn’t need to know concrete types.

Characteristics:

Centralized creation logic.

Can use switch or map internally.

Often called a “Simple Factory” in textbooks (not officially a Go pattern, just a helper function).

func PaymentFactory(method string) PaymentMethod {
    switch method {
    case "paypal":
        return &PayPal{}
    case "credit":
        return &CreditCard{}
    }
    return nil
}

Tight coupling to concrete types

The factory function knows every concrete implementation.

If you add upi, you must modify this factory.

Violates OCP (Open/Closed Principle)

Code is not closed for modification. Each new type requires touching the switch.

Runtime decisions

The mapping between method → type is resolved at runtime, which can lead to errors if a name is mistyped or missing.

Factory Method Pattern 

Advantages of This Approach

No switch or map → fully decoupled.

OCP-friendly → Add new notifier + factory without touching existing code.

Polymorphism → Client code depends only on NotifierFactory interface.

SOLID-friendly:

SRP → Each factory creates one product type.

OCP → Can add new notifiers easily.

DIP → Client depends on abstractions (NotifierFactory), not concrete types.

Using Map 

Case A : You modify the registry directly
var registry = map[string]func() Notifier{
    "email": func() Notifier { return &EmailNotifier{} },
    "sms":   func() Notifier { return &SMSNotifier{} },
    "push":  func() Notifier { return &PushNotifier{} }, // added here
}

This does violate OCP, because you are modifying the existing registry (which is essentially the factory). You touched already tested code.

Case B: You extend by calling RegisterNotifier in a new package/file
func init() {
    RegisterNotifier("push", func() Notifier { return &PushNotifier{} })
}

This respects OCP:

You didn’t modify the factory.

You simply extended the system by registering a new type at runtime.



