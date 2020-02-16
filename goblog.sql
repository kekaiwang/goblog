/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : blog_sql

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 16/02/2020 19:24:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wkk_admin_user
-- ----------------------------
DROP TABLE IF EXISTS `wkk_admin_user`;
CREATE TABLE `wkk_admin_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `login_count` int(11) DEFAULT '0' COMMENT '登陆次数',
  `salt` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL COMMENT '1:正常 2:禁止登陆',
  `last_login` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次登陆',
  `last_ip` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wkk_admin_user
-- ----------------------------
BEGIN;
INSERT INTO `wkk_admin_user` VALUES (1, 'admin', 'b8700fbb63d18d29164285cbb996c5f3', 'wkekai@163.com', 92, '$&)w', 1, '2020-02-13 16:39:18', '127.0.0.1', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for wkk_article
-- ----------------------------
DROP TABLE IF EXISTS `wkk_article`;
CREATE TABLE `wkk_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `count` int(11) unsigned DEFAULT '0' COMMENT '评论数量',
  `markdown` text CHARACTER SET utf8mb4,
  `content` text CHARACTER SET utf8mb4 COMMENT '内容',
  `category_id` int(11) DEFAULT NULL COMMENT '归属专题',
  `tag_ids` varchar(255) DEFAULT NULL COMMENT '标签分类',
  `excerpt` varchar(255) DEFAULT NULL COMMENT '预览信息',
  `previous` varchar(255) DEFAULT NULL COMMENT '前一篇',
  `next` varchar(255) DEFAULT NULL COMMENT '后一篇',
  `preview` int(11) DEFAULT NULL COMMENT '浏览数量',
  `thumb` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `slug` varchar(255) NOT NULL COMMENT '路由地址',
  `is_draft` tinyint(1) DEFAULT '1' COMMENT '1：草稿 2:已发布 3:已删除',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `edit_time` datetime DEFAULT NULL COMMENT '编辑时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `display_time` datetime DEFAULT NULL COMMENT '显示时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wkk_article
-- ----------------------------
BEGIN;
INSERT INTO `wkk_article` VALUES (1, 'kk wang', '开篇自述', 0, '#### 开始\n终于开始了好久的想法，blog - 又称博客。\n作为自己记录与激励和分享的一种方式，同时也是和大家一起交流学习的途径。\n\n#### 写在前面 ####\n> 伟人曾说过：好的开始就是成功的一半，so磕磕啪啪利用不到两周的时间使用Golang完成了首个自己的重要项目。\n\n最近赋闲在家，记录下自己的菜X程序生涯，半路出家到现在做开发已有四年的时间，经历的第一次离职，自己想着写点小东西，也是四年前的想法-blog，终于咻咻咻咻的开始了。。。\n#### 关于本站\n> 愿你出走半生归来仍是少年。\n\n基础架构：Golang-Beego + Mysql + Vue + Nginx\n现有List、Category、Tag、About和Links，后期再开通归档和评论。大部分保持原创分享同时不定期更新Links用于分享好的站点和知识。\n> 可能有的同学会问那么多好的开源blog不用非要自己写？也是为了能够有一个很好的学习go的机会和开源分享的机会，怎能放过？\n\n#### 简单介绍下自己\n本人姓王，名kekai。2012世界末日这年大学毕业，中间`辗转反侧`机缘巧合的情况下来到了北京加入了程序员的行列，要感谢老婆和家人的大力支持。\nPHP开发到现在，曾任职于一家”美股上市“公司做部门主力搬砖人员。\n熟练搬运PHP代码、Html、Js代码、Vue（Element、Iview）砖块，最近对Golang有一些了解和使用。\n\n希望通过自己的分享与学习可以认识更多朋友向大神多多学习！\n\n### **Find me**\n- email：[wkekai@163.com](mailto:wkekai@163.com)\n- github：年前后整理好代码后再来update\n\n```language\n注：\n\n构思途中参考了（排名不分先后）：\n- [deepzz](https://deepzz.com/)\n- [XIN.](https://lscho.com/)\n- [Razeen](https://razeencheng.com/)\n\n开源不易，大力支持！\n```\n', '<h4><a id=\"_0\"></a>开始</h4>\n<p>终于开始了好久的想法，blog - 又称博客。<br />\n作为自己记录与激励和分享的一种方式，同时也是和大家一起交流学习的途径。</p>\n<h4><a id=\"_4\"></a>写在前面</h4>\n<blockquote>\n<p>伟人曾说过：好的开始就是成功的一半，so磕磕啪啪利用不到两周的时间使用Golang完成了首个自己的重要项目。</p>\n</blockquote>\n<p>最近赋闲在家，记录下自己的菜X程序生涯，半路出家到现在做开发已有四年的时间，经历的第一次离职，自己想着写点小东西，也是四年前的想法-blog，终于咻咻咻咻的开始了。。。</p>\n<h4><a id=\"_8\"></a>关于本站</h4>\n<blockquote>\n<p>愿你出走半生归来仍是少年。</p>\n</blockquote>\n<p>基础架构：Golang-Beego + Mysql + Vue + Nginx<br />\n现有List、Category、Tag、About和Links，后期再开通归档和评论。大部分保持原创分享同时不定期更新Links用于分享好的站点和知识。</p>\n<blockquote>\n<p>可能有的同学会问那么多好的开源blog不用非要自己写？也是为了能够有一个很好的学习go的机会和开源分享的机会，怎能放过？</p>\n</blockquote>\n<h4><a id=\"_15\"></a>简单介绍下自己</h4>\n<p>本人姓王，名kekai。2012世界末日这年大学毕业，中间<code>辗转反侧</code>机缘巧合的情况下来到了北京加入了程序员的行列，要感谢老婆和家人的大力支持。<br />\nPHP开发到现在，曾任职于一家”美股上市“公司做部门主力搬砖人员。<br />\n熟练搬运PHP代码、Html、Js代码、Vue（Element、Iview）砖块，最近对Golang有一些了解和使用。</p>\n<p>希望通过自己的分享与学习可以认识更多朋友向大神多多学习！</p>\n<h3><a id=\"Find_me_22\"></a><strong>Find me</strong></h3>\n<ul>\n<li>email：<a href=\"mailto:wkekai@163.com\" target=\"_blank\">wkekai@163.com</a></li>\n<li>github：年前后整理好代码后再来update</li>\n</ul>\n<pre><code class=\"lang-language\">注：\n\n构思途中参考了（排名不分先后）：\n- [deepzz](https://deepzz.com/)\n- [XIN.](https://lscho.com/)\n- [Razeen](https://razeencheng.com/)\n\n开源不易，大力支持！\n</code></pre>\n', 4, '3,4', 'Wangkekai‘s blog 开篇自述，作为自己记录与激励的一种方式，同时也是和大家一起交流学习的一种途径。', '', 'service-config.html', 112, '', 'first-start.html', 2, '2020-01-08 20:43:01', NULL, '2020-01-15 23:01:54', '2020-01-08 20:31:59');
COMMIT;

-- ----------------------------
-- Table structure for wkk_article_relation
-- ----------------------------
DROP TABLE IF EXISTS `wkk_article_relation`;
CREATE TABLE `wkk_article_relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) DEFAULT NULL COMMENT '文章id',
  `tag_id` int(11) DEFAULT NULL COMMENT '标签id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wkk_article_relation
-- ----------------------------
BEGIN;
INSERT INTO `wkk_article_relation` VALUES (1, 1, 3);
INSERT INTO `wkk_article_relation` VALUES (2, 1, 4);
COMMIT;

-- ----------------------------
-- Table structure for wkk_category
-- ----------------------------
DROP TABLE IF EXISTS `wkk_category`;
CREATE TABLE `wkk_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '分类名字',
  `router_link` varchar(255) DEFAULT NULL COMMENT '路由地址',
  `link_article` int(11) unsigned DEFAULT '0' COMMENT '关联文章数量',
  `status` tinyint(1) DEFAULT NULL COMMENT '1:正常 2:禁用 3:已删除',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='文章分类表';

-- ----------------------------
-- Records of wkk_category
-- ----------------------------
BEGIN;
INSERT INTO `wkk_category` VALUES (1, '日常总结', 'daySummary.html', 2, 1, '2020-01-08 20:30:57', '2020-01-11 17:20:11');
INSERT INTO `wkk_category` VALUES (2, 'PHP', 'php.html', 0, 2, '2020-01-08 20:31:04', '2020-01-11 17:20:02');
INSERT INTO `wkk_category` VALUES (3, 'Golang', 'golang.html', 0, 1, '2020-01-08 20:31:12', '2020-01-11 17:19:55');
INSERT INTO `wkk_category` VALUES (4, '我的故事', 'myLife.html', 0, 1, '2020-01-08 20:31:26', '2020-01-11 17:19:42');
COMMIT;

-- ----------------------------
-- Table structure for wkk_page_info
-- ----------------------------
DROP TABLE IF EXISTS `wkk_page_info`;
CREATE TABLE `wkk_page_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '页面名字',
  `slug` varchar(255) DEFAULT NULL COMMENT '页面标识,路由地址',
  `content` text COMMENT '内容',
  `markdown` text,
  `status` tinyint(1) DEFAULT '1' COMMENT '1:正常 2:禁用 3:已删除',
  `preview` int(11) DEFAULT '0' COMMENT '阅读数量',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wkk_page_info
-- ----------------------------
BEGIN;
INSERT INTO `wkk_page_info` VALUES (1, '关于', 'about', '<h3><a id=\"_0\"></a><strong>关于我</strong></h3>\n<p>本人姓王，名<code>·kekai·</code>，籍贯：水泊梁山。2012年毕业于美丽的夏都西宁<code>昆院</code>，2015年兜兜溜溜(机缘巧合)来到首都开始了码农生涯。<br />\n现居北京,曾任职于一家美股上市公司做主力PHP搬砖人员。</p>\n<h3><a id=\"_3\"></a><strong>一些属性</strong></h3>\n<ul>\n<li>PHP开发者、Golang爱好者、Vue开发者、web全站开发；</li>\n<li>Mysql、Redis；</li>\n<li>Git、SVN版本控制；</li>\n<li>热爱篮球、跑步、旅行(浪)、美食、来到首都后喜欢上了厨房。</li>\n</ul>\n<h3><a id=\"_8\"></a><strong>假如或许想要联系我</strong></h3>\n<ul>\n<li>email：<a href=\"mailto:wkekai@163.com\" target=\"_blank\">wkekai@163.com</a></li>\n<li>github：年前后整理好代码后再来update</li>\n</ul>\n<h3><a id=\"_11\"></a><strong>关于本站</strong></h3>\n<ul>\n<li>Golang-Beego + Mysql + Nginx + Qiniu 存储</li>\n<li>服务器使用阿里云 CentOS 1G + 1M + 40G</li>\n<li>管理端使用Vue-Element UI + Mavon editor编辑文章</li>\n<li>待内容稍丰后评论、搜索等后续开通</li>\n<li>初衷为了实践Golang的实际使用和学习，后面会在这里作为自己的一种记录和分享<br />\n不仅限于知识技术的记录分享也会有自己的生活和其他的一些东东。。</li>\n</ul>\n<blockquote>\n<p>本站前端UI借鉴使用了「 <a href=\"https://lscho.com/\" target=\"_blank\">XIN.</a> 」的样式，管理端使用Vue完成。</p>\n</blockquote>\n<p>再聊聊logo吧，因为我家姑娘小名单字<code>淼</code>，自己喜欢叫自己<code>王三水</code>，so…logo就这么来了。</p>\n<h3><a id=\"_23\"></a><strong>版权说明</strong></h3>\n<p>本博客所有文章除特别声明外，均采用「 <a href=\"https://creativecommons.org/licenses/by/4.0/\" target=\"_blank\">知识共享署名4.0</a> 」创作共享协议。通俗地讲，只要在使用时署名，那么使用者可以对本站所有原创内容进行转载、节选、混编、二次创作，允许商业性使用。如有其他侵权请联系删除。</p>\n', '### **关于我**\n本人姓王，名`·kekai·`，籍贯：水泊梁山。2012年毕业于美丽的夏都西宁`昆院`，2015年兜兜溜溜(机缘巧合)来到首都开始了码农生涯。\n现居北京,曾任职于一家美股上市公司做主力PHP搬砖人员。\n### **一些属性**\n- PHP开发者、Golang爱好者、Vue开发者、web全站开发；\n- Mysql、Redis；\n- Git、SVN版本控制；\n- 热爱篮球、跑步、旅行(浪)、美食、来到首都后喜欢上了厨房。\n### **假如或许想要联系我**\n- email：[wkekai@163.com](mailto:wkekai@163.com)\n- github：年前后整理好代码后再来update\n### **关于本站**\n- Golang-Beego + Mysql + Nginx + Qiniu 存储\n- 服务器使用阿里云 CentOS 1G + 1M + 40G\n- 管理端使用Vue-Element UI + Mavon editor编辑文章\n- 待内容稍丰后评论、搜索等后续开通\n- 初衷为了实践Golang的实际使用和学习，后面会在这里作为自己的一种记录和分享\n  不仅限于知识技术的记录分享也会有自己的生活和其他的一些东东。。\n\n> 本站前端UI借鉴使用了「 [XIN.](https://lscho.com/) 」的样式，管理端使用Vue完成。\n\n再聊聊logo吧，因为我家姑娘小名单字`淼`，自己喜欢叫自己`王三水`，so...logo就这么来了。\n\n### **版权说明**\n本博客所有文章除特别声明外，均采用「 [知识共享署名4.0](https://creativecommons.org/licenses/by/4.0/) 」创作共享协议。通俗地讲，只要在使用时署名，那么使用者可以对本站所有原创内容进行转载、节选、混编、二次创作，允许商业性使用。如有其他侵权请联系删除。\n', 1, 0, '2020-01-10 22:21:13', '2020-02-08 21:56:56');
INSERT INTO `wkk_page_info` VALUES (2, '链接', 'links', '<h4><a id=\"_0\"></a>一些链接</h4>\n<p><a href=\"https://deepzz.com/\" target=\"_blank\">Deepzz</a> 「 唯爱与美食不可辜负也 」<br />\n<a href=\"https://razeencheng.com/\" target=\"_blank\">Razeen</a> 「 Stay hungry, Stay foolish. 」<br />\n<a href=\"https://lscho.com/\" target=\"_blank\">XIN.</a></p>\n<pre><code class=\"lang-language\">构思开发中借鉴的诸位大佬先列上(排名不分先后)\n</code></pre>\n<p><a href=\"https://xueyuanjun.com/\" target=\"_blank\">学院君(原Laravel学院)</a> 「 PHP Laravel框架的学习胜地 」<br />\n<a href=\"https://beego.me/\" target=\"_blank\">Beego</a>「 不知道怎么选择时相信自己人没有错 」<br />\n<a href=\"https://ustbhuangyi.github.io/vue-analysis/\" target=\"_blank\">vue技术揭秘</a> 「 vue原理的讲解非常好 」<br />\n<a href=\"https://github.com/PanJiaChen/vue-element-admin\" target=\"_blank\">Element-UI</a> 「 在公司的好几个项目上用过，so自己的管理端也是它了 」</p>\n<blockquote>\n<p>会不定期更新自己感觉好的知识或者链接地址</p>\n</blockquote>\n', '#### 一些链接\n[Deepzz](https://deepzz.com/) 「 唯爱与美食不可辜负也 」\n[Razeen](https://razeencheng.com/) 「 Stay hungry, Stay foolish. 」\n[XIN.](https://lscho.com/) \n\n```language\n构思开发中借鉴的诸位大佬先列上(排名不分先后)\n```\n\n[学院君(原Laravel学院)](https://xueyuanjun.com/) 「 PHP Laravel框架的学习胜地 」\n[Beego](https://beego.me/)「 不知道怎么选择时相信自己人没有错 」\n[vue技术揭秘](https://ustbhuangyi.github.io/vue-analysis/) 「 vue原理的讲解非常好 」\n[Element-UI](https://github.com/PanJiaChen/vue-element-admin) 「 在公司的好几个项目上用过，so自己的管理端也是它了 」\n\n> 会不定期更新自己感觉好的知识或者链接地址\n', 1, 0, '2020-01-10 22:42:07', '2020-01-14 21:44:22');
COMMIT;

-- ----------------------------
-- Table structure for wkk_request
-- ----------------------------
DROP TABLE IF EXISTS `wkk_request`;
CREATE TABLE `wkk_request` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `refer` varchar(255) DEFAULT NULL COMMENT '请求来源',
  `url` varchar(255) DEFAULT NULL COMMENT '访问页面',
  `major` int(11) DEFAULT NULL COMMENT '主版本',
  `ip` varchar(255) DEFAULT NULL COMMENT '请求IP',
  `proxy` varchar(255) DEFAULT NULL COMMENT '代理地址',
  `session_id` varchar(255) DEFAULT NULL COMMENT '请求session',
  `user_agent` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for wkk_tag
-- ----------------------------
DROP TABLE IF EXISTS `wkk_tag`;
CREATE TABLE `wkk_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `router_link` varchar(255) DEFAULT NULL COMMENT '路由链接',
  `use_times` int(11) DEFAULT '0' COMMENT '使用次数',
  `status` tinyint(1) DEFAULT '1' COMMENT '1:正常 2:禁用',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='标签数据';

-- ----------------------------
-- Records of wkk_tag
-- ----------------------------
BEGIN;
INSERT INTO `wkk_tag` VALUES (1, 'PHP', 'php.html', 1, 1, '2020-01-08 20:26:52', '2020-01-11 17:23:16');
INSERT INTO `wkk_tag` VALUES (2, 'JS', 'js.html', 1, 1, '2020-01-08 20:26:56', '2020-01-11 17:23:11');
INSERT INTO `wkk_tag` VALUES (3, 'my life', 'myLife.html', 1, 1, '2020-01-08 20:27:03', '2020-01-11 17:23:03');
INSERT INTO `wkk_tag` VALUES (4, 'Golang', 'golang.html', 2, 1, '2020-01-08 20:27:08', '2020-01-11 17:22:58');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
