# go-bank-service

## Overview
`go-bank-service` เป็นโปรเจ็กต์ฝึกในการสร้าง API service โดยใช้ภาษา Go โดยในโปรเจ็กต์นี้จะมีการใช้ [Fiber](https://gofiber.io/) ซึ่งเป็น web framework , [GORM](https://gorm.io/) ซึ่งเป็น ORM (Object-Relational Mapping) , ใช้และโครงสร้างแบบ Hexagonal Architecture เพื่อทำให้โค้ดมีการจัดการที่ดีและง่ายต่อการบำรุงรักษา

## Features
- สร้างและจัดการบัญชีธนาคาร
- ฝากและถอนเงิน
- ดูประวัติการทำธุรกรรม
- ใช้ Fiber ในการสร้าง RESTful API
- ใช้ GORM ในการเชื่อมต่อและจัดการฐานข้อมูล
- ออกแบบโครงสร้างโปรเจ็กต์ตามแนวคิด Hexagonal Architecture

## Requirements
- Go 1.18+
- Docker (สำหรับการตั้งค่าและใช้งานฐานข้อมูล)
- make (สำหรับคำสั่ง build และ run)

## Getting Started
### การติดตั้งและรันโปรเจ็กต์
1. คลอนโค้ดโปรเจ็กต์จาก GitHub:
   ```sh
   git clone https://github.com/PParist/go-bank-service.git
   cd go-bank-service
