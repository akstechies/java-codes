package com.akstechies;

public class Varargs {
    
    // Variable Arguments (Varargs) in Java is a method that takes a variable number of arguments. Variable Arguments in Java simplify the creation of methods that need to take a variable number of arguments.
    static void Names(String... n) {
        for (String i : n) {
            System.out.println(i);
        }
    }

}
