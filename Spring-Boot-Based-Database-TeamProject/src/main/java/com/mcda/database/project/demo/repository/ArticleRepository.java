package com.mcda.database.project.demo.repository;

import com.mcda.database.project.demo.model.Article;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;

import java.util.List;
import java.util.Optional;

public interface ArticleRepository extends JpaRepository<Article, Integer> {

    List<Article> findByMagazineIdAndVolume(String magazineId, String volume);


    Optional<Article> findByVolumeNumber(int volumeNumber);
}
