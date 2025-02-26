package com.akstechies;

import java.util.Arrays;

public class JavaArrays {
    static void declare_arrays() {
        
        int[] arr = {1, 2, 3, 4, 5, 6};
        System.out.println(Arrays.toString(arr));
        System.out.println(arr.length);

        // declare array of integers
        int[] arr2;

        // allocate space for arrays
        arr2 = new int[5];

        // give them the values
        arr2[0] = 1;
        arr2[1] = 2;
        arr2[2] = 3;
        arr2[3] = 4;
        arr2[4] = 5;
        System.out.println(Arrays.toString(arr2));


        // object as arrays
    }
}

class Student {
    public int roll_no;
    public String name;

    Student(int roll_no, String name) {
        this.roll_no = roll_no;
        this.name = name;
    }

}
