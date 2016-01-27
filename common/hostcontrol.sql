-- MySQL dump 10.16  Distrib 10.1.10-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: hostcontrol
-- ------------------------------------------------------
-- Server version	10.1.10-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `domainmetadata`
--

DROP TABLE IF EXISTS `domainmetadata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `domainmetadata` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) DEFAULT NULL,
  `kind` varchar(32) DEFAULT NULL,
  `content` text,
  PRIMARY KEY (`id`),
  KEY `domainidmetaindex` (`domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `domains`
--

DROP TABLE IF EXISTS `domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `domains` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `master` varchar(128) DEFAULT NULL,
  `last_check` int(11) DEFAULT NULL,
  `type` varchar(6) NOT NULL,
  `notified_serial` int(11) DEFAULT NULL,
  `account` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hostcontrol_ftpusers`
--

DROP TABLE IF EXISTS `hostcontrol_ftpusers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hostcontrol_ftpusers` (
  `ftpuser_id` int(11) NOT NULL AUTO_INCREMENT,
  `ftpusername` varchar(32) NOT NULL,
  `homedir` varchar(255) NOT NULL,
  `system_username` varchar(32) NOT NULL,
  PRIMARY KEY (`ftpuser_id`),
  KEY `ftpusername` (`ftpusername`,`system_username`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hostcontrol_ipaddr`
--

DROP TABLE IF EXISTS `hostcontrol_ipaddr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hostcontrol_ipaddr` (
  `ip_id` int(11) NOT NULL AUTO_INCREMENT,
  `system_username` varchar(32) NOT NULL,
  `ipaddr` varchar(16) NOT NULL,
  PRIMARY KEY (`ip_id`),
  UNIQUE KEY `system_username_2` (`system_username`),
  KEY `system_username` (`system_username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hostcontrol_user_tokens`
--

DROP TABLE IF EXISTS `hostcontrol_user_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hostcontrol_user_tokens` (
  `token` varchar(34) NOT NULL,
  `hostcontrol_id` int(11) NOT NULL,
  `description` varchar(45) NOT NULL,
  `token_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`token_id`),
  KEY `hostcontrol_id` (`hostcontrol_id`),
  CONSTRAINT `deleteuser` FOREIGN KEY (`hostcontrol_id`) REFERENCES `hostcontrol_users` (`hostcontrol_id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hostcontrol_users`
--

DROP TABLE IF EXISTS `hostcontrol_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hostcontrol_users` (
  `hostcontrol_id` int(11) NOT NULL AUTO_INCREMENT,
  `system_username` varchar(32) NOT NULL,
  `privileges` text NOT NULL,
  `owned_by` varchar(32) NOT NULL,
  `login_token` varchar(32) NOT NULL,
  `email_address` varchar(255) NOT NULL,
  PRIMARY KEY (`hostcontrol_id`),
  KEY `owned_by` (`owned_by`),
  KEY `system_username` (`system_username`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_domains`
--

DROP TABLE IF EXISTS `mail_domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_domains` (
  `domain_id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) NOT NULL,
  `system_username` varchar(32) NOT NULL,
  PRIMARY KEY (`domain_id`),
  UNIQUE KEY `domain_2` (`domain`),
  KEY `domain` (`domain`),
  KEY `system_username` (`system_username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_forwardings`
--

DROP TABLE IF EXISTS `mail_forwardings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_forwardings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source` varchar(80) NOT NULL,
  `destination` text NOT NULL,
  `domain` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `domain` (`domain`),
  KEY `source` (`source`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_transport`
--

DROP TABLE IF EXISTS `mail_transport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_transport` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(128) NOT NULL DEFAULT '',
  `transport` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain` (`domain`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_users`
--

DROP TABLE IF EXISTS `mail_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_users` (
  `email_id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(20) NOT NULL,
  `domain` varchar(255) NOT NULL,
  PRIMARY KEY (`email_id`),
  KEY `domain` (`domain`),
  KEY `email` (`email`),
  CONSTRAINT `mail_users_ibfk_1` FOREIGN KEY (`domain`) REFERENCES `mail_domains` (`domain`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `records`
--

DROP TABLE IF EXISTS `records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  `content` varchar(64000) DEFAULT NULL,
  `ttl` int(11) DEFAULT NULL,
  `prio` int(11) DEFAULT NULL,
  `change_date` int(11) DEFAULT NULL,
  `disabled` tinyint(1) DEFAULT '0',
  `ordername` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `auth` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `nametype_index` (`name`,`type`),
  KEY `domain_id` (`domain_id`),
  KEY `recordorder` (`domain_id`,`ordername`),
  CONSTRAINT `records_ibfk_1` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `website_vhosts`
--

DROP TABLE IF EXISTS `website_vhosts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `website_vhosts` (
  `vhost_id` int(11) NOT NULL AUTO_INCREMENT,
  `system_username` varchar(32) NOT NULL,
  `domain` varchar(254) NOT NULL,
  `documentroot` text NOT NULL,
  `ipaddr` varchar(16) NOT NULL,
  `ssl_enabled` enum('Y','N') NOT NULL,
  `ssl_certificate` longtext NOT NULL,
  `ssl_key` longtext NOT NULL,
  `ssl_ca_certificate` longtext NOT NULL,
  PRIMARY KEY (`vhost_id`),
  KEY `domain` (`domain`),
  KEY `system_username` (`system_username`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-01-27 23:34:30
