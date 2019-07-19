/*
SQLyog Ultimate v12.08 (64 bit)
MySQL - 8.0.11 : Database - iris
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
USE `iris-admin`;

/*Table structure for table `admin` */

DROP TABLE IF EXISTS `admin`;

CREATE TABLE `admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `account` varchar(20) NOT NULL,
  `password` char(32) NOT NULL,
  `descript` varchar(255) DEFAULT '',
  `nickname` char(100) DEFAULT '',
  `email` varchar(100) DEFAULT '',
  `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `username` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `admin` */

insert  into `admin`(`id`,`deleted_at`,`account`,`password`,`descript`,`nickname`,`email`,`avatar`,`created_at`,`updated_at`) values (1,NULL,'admin','21232f297a57a5a743894a0e4a801fc3','一个golang iris学习者','灯火阑珊','7146275@qq.com','/uploads/avatar/3348_76c19bb79eba76d39c1b56c020571e17.jpg','2018-10-22 14:03:48','2019-04-08 11:07:17'),(2,NULL,'cuijun','3b7fb9742017f12726bcebcd69fb7470','Go Web Iris中文网致力于，在中国国内推广Go语言','众里寻他','10000@qq.cm','/uploads/avatar/4862_head_1 (4).png','2018-11-01 10:53:53','2018-11-01 14:15:55'),(3,NULL,'test','098f6bcd4621d373cade4e832627b4f6','testtesttesttest','test1231','test03@qq.com','/uploads/avatar/4548_9E4ACEB09C8EFDB2A08D8009EC616CB84304578F_size37_w550_h393.jpeg','2018-11-01 14:16:42','2018-12-21 18:13:13'),(4,NULL,'q3123123','e10adc3949ba59abbe56e057f20f883e','sdsaaaaaaaaaaaaaa','wqewqeqw','admin@dwqeq.com','/uploads/avatar/5239_192654486.jpg','2019-06-06 16:13:54','2019-06-06 16:13:54');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类表';

/*Data for the table `category` */

insert  into `category`(`id`,`name`,`parent_id`,`sort`,`created_at`,`updated_at`,`deleted_at`) values (1,'国家分类',0,999,'2018-11-06 10:00:29','2018-11-06 10:00:32',NULL),(2,'亚洲',1,0,'2018-11-06 10:00:45','2018-11-06 10:00:48',NULL),(3,'中国',2,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(4,'韩国',2,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(5,'日本',2,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(6,'北美洲',1,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(7,'欧洲',1,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(8,'南美洲',1,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(9,'非洲',0,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(10,'大洋洲',1,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(11,'美国',6,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(12,'加拿大',6,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(13,'墨西哥',6,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(14,'英国',7,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(15,'法国',7,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(16,'德国',7,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(17,'巴西',8,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(18,'阿根廷',8,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(19,'秘鲁',8,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(20,'埃及',9,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(21,'南非',9,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(22,'肯尼亚',9,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(23,'澳大利亚',10,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL),(24,'新西兰',10,0,'2018-11-06 10:00:45','2018-11-06 10:00:45',NULL);

/*Table structure for table `menu` */

DROP TABLE IF EXISTS `menu`;

CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `menu` */

insert  into `menu`(`id`,`name`,`parent_id`,`sort`,`created_at`,`updated_at`,`deleted_at`) values (1,'测试5',0,0,'0000-00-00 00:00:00','2019-06-06 13:58:56',NULL);

/*Table structure for table `news` */

DROP TABLE IF EXISTS `news`;

CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` varchar(100) NOT NULL,
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '分类名称',
  `descript` varchar(500) NOT NULL DEFAULT '' COMMENT '父id',
  `content` text NOT NULL,
  `tags` varchar(100) DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `sort` (`sort`),
  KEY `category_id` (`category_id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='内容表';

/*Data for the table `news` */

insert  into `news`(`id`,`category_id`,`title`,`descript`,`content`,`tags`,`sort`,`created_at`,`updated_at`,`deleted_at`) values (1,'2,3','测试1','测试1','<p>测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1测试1</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><figure class=\"table\"><table><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table></figure>','',1,'0000-00-00 00:00:00','2019-06-06 15:06:09',NULL),(2,'3','213','wd','<p><strong>dwdwdwd</strong></p>\r\n','',1,'2018-11-07 15:35:33','2018-11-07 15:35:33',NULL),(3,'3','324234','efef','<p>ef</p>\r\n','',1,'2018-11-07 15:36:20','2018-11-07 15:36:20',NULL),(4,'2,4,6,20,22','234324','我的','<p>二次沟</p>\r\n\r\n<table border=\"1\" cellpadding=\"1\" cellspacing=\"1\" style=\"width:500px\">\r\n	<tbody>\r\n		<tr>\r\n			<td>\r\n			<h2>多吃点</h2>\r\n			</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n		<tr>\r\n			<td>&nbsp;</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n		<tr>\r\n			<td>&nbsp;</td>\r\n			<td>&nbsp;</td>\r\n		</tr>\r\n	</tbody>\r\n</table>\r\n\r\n<p>&nbsp;</p>\r\n','',1,'2018-11-07 15:42:02','2018-11-07 16:36:23',NULL),(5,'23','存储','萨达飒飒','<figure class=\"table\"><table><tbody><tr><td>123</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table></figure><p>&nbsp;</p><p>&nbsp;</p><p>sdasdas</p><p>&nbsp;</p><p>&nbsp;</p><p>dfsdf</p>','',0,'0000-00-00 00:00:00','2019-06-06 11:36:05',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
