package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.TransactionItems;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;

public interface TransactionItemsRepository extends JpaRepository<TransactionItems, Integer> {
    void deleteByTransactionNumber(int transactionNumber);
}
