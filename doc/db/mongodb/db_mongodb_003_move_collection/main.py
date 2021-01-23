#!/usr/bin/python
# -*- coding: UTF-8 -*-
import pymongo


def move():
    client = pymongo.MongoClient("mongodb://localhost:27017")
    old_items = client["repeat"]["item"].find()
    new_col = client["boat"]["review_item"]
    new_items = []

    for item in old_items:
        new_items.append(item)
    new_col.insert_many(new_items)


if __name__ == "__main__":
    move()
