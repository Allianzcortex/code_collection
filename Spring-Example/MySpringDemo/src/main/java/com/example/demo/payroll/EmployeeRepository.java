package com.example.demo.payroll;

import org.springframework.data.jpa.repository.JpaRepository;

/**
 * 虽然这个时候 interface 仍然是空的，但已经可以支持
 * Creating new instances

Updating existing ones

Deleting

Finding (one, all, by simple or complex properties)
 */
interface EmployeeRepository extends JpaRepository<Employee,Long>{

}