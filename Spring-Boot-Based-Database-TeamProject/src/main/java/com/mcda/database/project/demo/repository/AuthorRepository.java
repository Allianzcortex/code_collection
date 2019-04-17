package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.Author;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;

public interface AuthorRepository extends JpaRepository<Author, Integer> {
}
