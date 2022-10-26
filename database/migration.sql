CREATE TABLE board
(
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(200),
    price INTEGER,
    descr VARCHAR(1000),
    date TIMESTAMP
);

CREATE TABLE photo
(
    post_id INTEGER,
    main BOOLEAN,
    link TEXT,
    FOREIGN KEY (post_id) REFERENCES board (post_id) ON DELETE CASCADE
);

INSERT INTO board (title,price,descr,date)
VALUES  ('car',1500,'bmx x3 2014',CURRENT_TIMESTAMP(2)),
        ('macbook',350,'M1 2021 256gb',CURRENT_TIMESTAMP(2)),
        ('table',60,'computer table',CURRENT_TIMESTAMP(2)),
        ('iphone',280,'iphone 6s',CURRENT_TIMESTAMP(2)),
        ('sneakers',75,'nike air force',CURRENT_TIMESTAMP(2)),
        ('TV',200,'LG smart TV',CURRENT_TIMESTAMP(2)),
        ('headphones',30,'honor am61',CURRENT_TIMESTAMP(2)),
        ('airpods',80,'airpods 2',CURRENT_TIMESTAMP(2)),
        ('teapot',55,'chineese teapot',CURRENT_TIMESTAMP(2)),
        ('jacket',100,'columbia winter jacket',CURRENT_TIMESTAMP(2)),
        ('car',20000,'toyota crown',CURRENT_TIMESTAMP(2)),
        ('xiaomi phone',120,'xiaomi mi5',CURRENT_TIMESTAMP(2)),
        ('huawei laptop',300,'huawei matepad',CURRENT_TIMESTAMP(2)),
        ('sneakers',65,'adidas streetball',CURRENT_TIMESTAMP(2)),
        ('sweater',50,'winter sweater',CURRENT_TIMESTAMP(2)),
        ('t-shirt',10,'kappa t-shirt',CURRENT_TIMESTAMP(2)),
        ('chair',70,'office chair',CURRENT_TIMESTAMP(2)),
        ('airpods',110,'airpods pro',CURRENT_TIMESTAMP(2));

INSERT INTO photo (post_id,main,link)
VALUES  (1,TRUE,'http://photo1'),
        (1,FALSE,'http://photo2'),
        (1,FALSE,'http://photo3'),
        (2,TRUE,'http://photo1'),
        (2,FALSE,'http://photo2'),
        (3,TRUE,'http://photo1'),
        (4,TRUE,'http://photo1'),
        (4,FALSE,'http://photo2'),
        (5,TRUE,'http://photo1'),
        (6,TRUE,'http://photo1'),
        (7,TRUE,'http://photo1'),
        (8,TRUE,'http://photo1'),
        (8,FALSE,'http://photo2'),
        (9,TRUE,'http://photo1'),
        (10,TRUE,'http://photo1'),
        (10,FALSE,'http://photo2'),
        (11,TRUE,'http://photo1'),
        (11,FALSE,'http://photo2'),
        (12,TRUE,'http://photo1'),
        (13,TRUE,'http://photo1'),
        (13,FALSE,'http://photo2'),
        (14,TRUE,'http://photo1'),
        (14,FALSE,'http://photo2'),
        (15,TRUE,'http://photo1'),
        (16,TRUE,'http://photo1'),
        (17,TRUE,'http://photo1'),
        (18,TRUE,'http://photo1'),
        (18,FALSE,'http://photo2');

