package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.Customer;
import org.springframework.data.repository.CrudRepository;

public interface CustomerRepository extends CrudRepository<Customer, Integer> {
    boolean existsByFirstNameAndLastName(String firstName, String lastName);


}
