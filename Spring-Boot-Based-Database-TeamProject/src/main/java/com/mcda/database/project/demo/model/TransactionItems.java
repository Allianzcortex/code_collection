package com.mcda.database.project.demo.model;

import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.format.annotation.DateTimeFormat;

import javax.persistence.*;

@Data
@Entity
@NoArgsConstructor
@Table(name = "transaction_items")
public class TransactionItems {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private int Id;

    @Column(name = "transaction_number")
    private int transactionNumber;

    @Column(name = "item_id")
    private int itemId;

    @Column(name = "item_price")
    private float price;
}
