package com.akstechies;

public class Greet {

    // this is a default access modifier so it will be availaible inside the same package only, if we use this in another package it leads to error
    void greet_them() {
        System.out.println("Hello World");
    }

    // creating a private access modifier so it will be availaible inside the same class only, not even the same package
    private static void secret_info() {
        System.out.println("COde is xx56");
    }

    static void get_secret_info() {
        secret_info();
    }


    // The protected access modifier is specified using the keyword protected. The methods or data members declared as protected are accessible within the same package or subclasses in different packages.
    protected void student_subject() {
        System.out.println("protected subject");
    }


    // public - accessible everywhere
    public void college() {
        System.out.println("SRMU");
    }


}
