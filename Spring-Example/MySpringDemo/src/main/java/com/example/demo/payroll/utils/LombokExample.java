package payroll.utils;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;


@ToString
public class LombokExample {
    @Getter
    @Setter
    private int age;
    private String name;

    public static void main(String[] args) {
        LombokExample ex = new LombokExample();
        System.out.println(ex);
    }

}


