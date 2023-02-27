use project;

CREATE TABLE `discuss`
(
    `Id`          int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `student_id`  int          DEFAULT NULL COMMENT '学生Id',
    `teacher_id`  int          DEFAULT NULL COMMENT '教师Id',
    `content`     varchar(255) DEFAULT NULL COMMENT '内容',
    `create_time` datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='讨论表';