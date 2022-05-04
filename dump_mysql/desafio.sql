-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: desafio
-- ------------------------------------------------------
-- Server version	5.7.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `families`
--

DROP TABLE IF EXISTS `families`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `families` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `relationship_id` int(11) DEFAULT NULL,
  `person_id` int(11) DEFAULT NULL,
  `person_family_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_family_relationship_idx` (`relationship_id`),
  KEY `fk_family_people_idx` (`person_id`),
  KEY `fk_family_person_idx` (`person_family_id`),
  CONSTRAINT `fk_family_people` FOREIGN KEY (`person_id`) REFERENCES `people` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_family_relationship` FOREIGN KEY (`relationship_id`) REFERENCES `relationships` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_person_family` FOREIGN KEY (`person_family_id`) REFERENCES `people` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `families`
--

LOCK TABLES `families` WRITE;
/*!40000 ALTER TABLE `families` DISABLE KEYS */;
INSERT INTO `families` VALUES (1,2,1,6),(2,6,1,11),(3,6,1,12),(4,5,2,3),(5,2,2,7),(6,2,2,8),(7,6,2,11),(8,6,2,12),(9,6,2,13),(10,5,3,2),(11,2,3,8),(12,2,3,7),(13,6,3,11),(14,6,3,12),(15,6,3,13),(16,5,4,5),(17,2,4,9),(18,6,4,13),(19,6,4,14),(20,5,5,4),(21,2,5,9),(22,6,5,13),(23,6,5,14),(24,1,6,1),(25,1,6,2),(26,5,6,5),(27,2,6,11),(28,2,6,12),(29,1,7,2),(30,1,7,3),(31,3,7,8),(32,2,7,11),(33,2,7,12),(34,5,7,6),(35,1,8,2),(36,1,8,3),(37,5,8,9),(38,2,8,13),(39,3,8,7),(40,1,9,4),(41,1,9,5),(42,5,9,8),(43,5,9,10),(44,2,9,13),(45,2,9,14),(46,5,10,9),(47,2,10,14),(48,4,11,1),(49,4,11,2),(50,1,11,5),(51,1,11,6),(52,3,11,12),(53,4,12,1),(54,4,12,2),(55,1,12,5),(56,1,12,6),(57,3,12,11),(58,4,13,2),(59,4,13,3),(60,4,13,4),(61,4,13,5),(62,1,13,8),(63,1,13,9),(64,3,13,14),(65,4,14,4),(66,4,14,5),(67,1,14,9),(68,1,14,10),(69,3,14,13);
/*!40000 ALTER TABLE `families` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genealogies`
--

DROP TABLE IF EXISTS `genealogies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `genealogies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `person_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_genealogies_family1_idx` (`person_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genealogies`
--

LOCK TABLES `genealogies` WRITE;
/*!40000 ALTER TABLE `genealogies` DISABLE KEYS */;
INSERT INTO `genealogies` VALUES (1,1),(2,2),(3,3),(4,4),(5,5),(6,6),(7,7),(8,8),(9,9),(10,10),(11,11),(12,12),(13,13),(14,14);
/*!40000 ALTER TABLE `genealogies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `people`
--

DROP TABLE IF EXISTS `people`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `people` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `people`
--

LOCK TABLES `people` WRITE;
/*!40000 ALTER TABLE `people` DISABLE KEYS */;
INSERT INTO `people` VALUES (1,'Sony'),(2,'Martin'),(3,'Anastasia'),(4,'Ellen'),(5,'Oprah'),(6,'Mike'),(7,'Phobe'),(8,'Ursula'),(9,'Eric'),(10,'Ariel'),(11,'Dunny'),(12,'Bruce'),(13,'Jacqueline'),(14,'Melody'),(18,'teste'),(21,''),(22,'');
/*!40000 ALTER TABLE `people` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `relationships`
--

DROP TABLE IF EXISTS `relationships`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `relationships` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relationships`
--

LOCK TABLES `relationships` WRITE;
/*!40000 ALTER TABLE `relationships` DISABLE KEYS */;
INSERT INTO `relationships` VALUES (1,'Parent'),(2,'Son'),(3,'Sibling'),(4,'Grandparent'),(5,'Life partner'),(6,'Grandson');
/*!40000 ALTER TABLE `relationships` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-04  2:09:55
