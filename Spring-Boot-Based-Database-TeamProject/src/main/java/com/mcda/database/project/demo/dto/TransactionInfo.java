package com.mcda.database.project.demo.dto;

import com.mcda.database.project.demo.model.Items;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
@NoArgsConstructor
public class TransactionInfo {

    private int customerId;

    private List<Items> items;
}
