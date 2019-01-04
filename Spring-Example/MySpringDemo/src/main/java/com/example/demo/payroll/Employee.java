package com.example.demo.payroll;

// Data 是 lombok ，帮助生成了 equals,hash,tostring 等方法
import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Data
@Entity  // 是 JPA 的注解，
public class Employee {
    private @Id
    @GeneratedValue
    Long id;
    private String name;
    private String role;

    Employee(String name, String role) {
        this.name = name;
        this.role = role;
    }

}
