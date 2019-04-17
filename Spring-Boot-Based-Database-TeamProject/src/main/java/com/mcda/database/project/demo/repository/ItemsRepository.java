package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.Items;
import org.springframework.data.repository.CrudRepository;


public interface ItemsRepository extends CrudRepository<Items,Integer> {

}
