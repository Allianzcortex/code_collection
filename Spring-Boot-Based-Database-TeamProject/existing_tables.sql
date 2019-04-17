DROP TABLE if exists AUTHOR;
#
create table if not exists AUTHOR (
  _id INT not null auto_increment,
  lname varchar(30) not null,
  fname varchar(30),
  email varchar(50),
  primary key(_id)
) engine = innodb;


DROP TABLE if exists MAGAZINE;
#
create table if not exists MAGAZINE (
  _id INT not null auto_increment,
  name varchar(50) not null,
  primary key(_id),
  check(name != '')
) engine = innodb;


DROP TABLE if exists ITEM;
#
create table if not exists ITEM (
  _id BIGINT not null auto_increment,
  price FLOAT(8,2) not null,
  primary key(_id)
) engine = innodb;


insert into MAGAZINE (name) values
("Theor. Comput. Sci."),
("Linguistics and Philosophy"),
("Lecture Notes in Computer Science");


insert into ITEM (price) values
(10.00),
(10.50),
(10.90),
(50.00),
(50.50),
(100.00),
(100.50),
(150.00);


insert into AUTHOR (fname,lname,email) values
("Houda", "Abbad", "houda.abbad@lycos.com"),
("Hideo", "Bannai", "bannai@inf.kyushu-u.ac.jp"),
("Mutlu", "Beyazit", "beyazit@adt.upb.de"),
("Francine", "Blanchet-Sadri", "blanchet@uncg.edu"),
("Janusz", "Brzozowski", "brzozo@uwaterloo.ca"),
("Cezar", "Campeanu", "cezar@sun11.math.upei.ca"),
("Mathieu", "Caralp", "mathieu.caralp@lif.univ-mrs.fr"),
("Pascal", "Caron", "pascal.caron@univ-rouen.fr"),
("Jean-Marc", "Champarnaud", "Jean-Marc.Champarnaud@univ-rouen.fr"),
("Dmitry", "Chistikov", "dch@mpi-sws.org"),
("Christian", "Choffrut", "Christian.Choffrut@liafa.jussieu.fr"),
("Stefano", "Crespi-Reghizzi", "stefano.crespireghizzi@polimi.it"),
("Denis", "Debarbieux", "denis.debarbieux@inria.fr"),
("Pierpaolo", "Degano", "degano@di.unipi.it"),
("Akim", "Demaille", "akim.demaille@gmail.com"),
("Michael", "Domaratzki", "mdomarat@cs.umanitoba.ca"),
("Frank", "Drewes", "drewes@cs.umu.se"),
("Alexandre", "Duret-Lutz", "adl@lrde.epita.fr"),
("Gianluigi", "Ferrari", "giangi@di.unipi.it"),
("Olivier", "Gauwin", "olivier.gauwin@labri.fr"),
("Thomas", "Genet", "thomas.genet@irisa.fr"),
("Daniela", "Genova", "d.genova@unf.edu"),
("Yuri", "Gurevich", "gurevich@microsoft.com"),
("Yo-Sub", "Han", "emmous@cs.yonsei.ac.kr"),
("MdMahbubul", "Hasan", "shanto86@gmail.com"),
("Pierre-Cyrille", "Heam", "pheam@femto-st.fr"),
("Fritz", "Henglein", "henglein@diku.dk"),
("Jan", "Holub", "Jan.Holub@fit.cvut.cz"),
("Markus", "Holzer", "holzer@informatik.uni-giessen.de"),
("Tomohiro", "I", "tomohiro.i@inf.kyushu-u.ac.jp"),
("A.S.M.Sohidull", "Islam", "sohansayed@gmail.com"),
("Masami", "Ito", "ito@cc.kyoto-su.ac.jp"),
("Sebastian", "Jakobi", "sebastian.jakobi@informatik.uni-giessen.de"),
("Jozef", "Jirasek", "jozef.jirasek@upjs.sk"),
("Oscar", "Ibarra", "ibarra@cs.ucsb.edu"),
("Shunsuke", "Inenaga", "inenaga@inf.kyushu-u.ac.jp"),
("Alois", "Dreyfus", "alois.dreyfus@femto-st.fr"),
("Galina", "Jiraskova", "jiraskov@saske.sk"),
("Natasa", "Jonoska", "jonoska@math.usf.edu"),
("Helmut", "Jurgensen", "hjj@csd.uwo.ca"),
("Lila", "Kari", "lila@csd.uwo.ca"),
("Andrzej", "Kisielewicz", "andrzej.kisielewicz@gmail.com"),
("Sang-Ki", "Ko", "narame7@cs.yonsei.ac.kr"),
("Stavros", "Konstantinidis", "s.konstantinidis@smu.ca"),
("Olga", "Kouchnarenko", "olga.kouchnarenko@femto-st.fr"),
("Dexter", "Kozen", "kozen@cs.cornell.edu"),
("Wener", "Kuich", "kuich@tuwien.ac.at"),
("Natalia", "Kushik", "ngkushik@gmail.com"),
("Martin", "Kutrib", "kutrib@informatik.uni-giessen.de"),
("Tristan", "LeGall", "tristan.le-gall@cea.fr"),
("Axel", "Legay", "axel.legay@inria.fr"),
("Pawan", "Lingras", "pawan.lingras@smu.ca"),
("Norma", "Linney", "norma.linney@smu.ca"),
("Sylvain", "Lombardy", "Sylvain.Lombardy@labri.fr"),
("Eva", "Maia", "emaia@dcc.fc.up.pt"),
("Rupak", "Majumdar", "rupak@mpi-sws.org"),
("Andreas", "Malcher", "andreas.malcher@informatik.uni-giessen.de"),
("Andreas", "Maletti", "andreas.maletti@ims.uni-stuttgart.de"),
("Sebastian", "Maneth", "Sebastian.Maneth@gmail.com"),
("Denis", "Maurel", "denis.maurel@univ-tours.fr"),
("Carlo", "Mereghetti", "mereghetti@di.unimi.it"),
("Gianluca", "Mezzetti", "mezzetti@di.unipi.it"),
("Nelma", "Moreira", "nam@dcc.fc.up.pt"),
("Frantisek", "Mraz", "mraz@ksvi.ms.mff.cuni.cz"),
("Paul", "Muir", "muir@smu.ca"),
("Valerie", "Murat", "valerie.murat@inria.fr"),
("Joachim", "Niehren", "joachim.niehren@inria.fr"),
("Lasse", "Nielsen", "lasse.nielsen.dk@gmail.com"),
("Takaaki", "Nishimoto", "a32b16c4@gmail.com"),
("Friedrich", "Otto", "otto@theory.informatik.uni-kassel.de"),
("Beatrice", "Palano", "palano@dsi.unimi.it"),
("Giovanni", "Pighizzini", "pighizzini@dico.unimi.it"),
("Daniel", "Prusa", "prusapa1@cmp.felk.cvut.cz"),
("M.Sohel", "Rahman", "sohel.kcl@gmail.com"),
("Ian", "McQuillan", "mcquillan@cs.usask.ca"),
("George", "Rahonis", "grahonis@math.auth.gr"),
("Bala", "Ravikumar", "ravi@cs.sonoma.edu"),
("Daniel", "Reidenbach", "D.Reidenbach@lboro.ac.uk"),
("Rogerio", "Reis", "rvr@dcc.fc.up.pt"),
("Pierre-Alain", "Reynier", "pierre-alain.reynier@lif.univ-mrs.fr"),
("Jacques", "Sakarovitch", "sakarovitch@telecom-paristech.fr"),
("Michel", "Rigo", "m.rigo@ulg.ac.be"),
("Kai", "Salomaa", "ksalomaa@cs.queensu.ca"),
("Pierluigi", "San-Pietro", "pierluigi.sanpietro@polimi.it"),
("Porter", "Scobey", "porter.scobey@smu.ca"),
("Tom", "Sebastian", "tom.sebastian@inria.fr"),
("Ayon", "Sen", "ayonsn@gmail.com"),
("Geraud", "Senizergues", "ges@labri.fr"),
("Klaus", "Sutner", "sutner@cs.cmu.edu"),
("Marek", "Szykula", "marek.esz@gmail.com"),
("Masayuki", "Takeda", "takeda@inf.kyushu-u.ac.jp"),
("Jean-Marc", "Talbot", "jean-marc.talbot@lif.univ-mrs.fr"),
("Marc", "Tommasi", "Marc.Tommasi@univ-lille3.fr"),
("Mikhail", "Volkov", "Mikhail.Volkov@usu.ru"),
("Bruce", "Watson", "bruce@fastar.org"),
("Matthias", "Wendlandt", "matthias.wendlandt@informatik.uni-giessen.de"),
("Hsu-Chun", "Yen", "yen@cc.ee.ntu.edu.tw"),
("Nina", "Yevtushenko", "ninayevtushenko@yahoo.com"),
("Mohamed", "Zergaoui", "innovimax@gmail.com"),
("Alexander", "Okhotin", "alexander.okhotin@utu.fi");
