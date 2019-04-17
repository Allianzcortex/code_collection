package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.Article;
import com.mcda.database.project.demo.model.ArticleAuthors;
import org.springframework.data.repository.CrudRepository;

public interface ArticleAuthorsRepository extends CrudRepository<ArticleAuthors,Integer> {
}
