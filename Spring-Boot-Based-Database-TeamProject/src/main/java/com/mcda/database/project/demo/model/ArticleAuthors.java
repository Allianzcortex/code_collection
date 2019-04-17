package com.mcda.database.project.demo.model;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.NonNull;
import lombok.RequiredArgsConstructor;

import javax.persistence.*;

@Data
@Entity
@NoArgsConstructor
@RequiredArgsConstructor
@Table(name = "article_authors")
public class ArticleAuthors {

    @javax.persistence.Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id")
    private int Id;

    @Column(name = "article_id")
    @NonNull
    private int articleId;

    @Column(name = "author_id")
    @NonNull
    private int authorId;
}
