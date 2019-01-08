package payroll.utils;

import lombok.ToString;

@ToString
public class LombokExample {
    private int age;
    private String name;

    public static void main(String[] args) {
        LombokExample ex = new LombokExample();
        System.out.println(ex);
    }

}


