# Data Transfer Utility
- [Redis](#redis)

## Redis
Redis is a shortened from *Remote Dictionary Server*. It stores a key-value data storage (that holds your data) in RAM Memory but does offer on-disk storage as its additional feature. <br>
Redis is capable of being a NoSQL database but it's more suitable for caching between your server code and your own database. <br>

<img src="https://hackajob.com/hubfs/Imported_Blog_Media/redis-1-3.png" width=450/><br>
*src: https://hackajob.com/hubfs/Imported_Blog_Media/redis-1-3.png*

### Best use case to implement cache
It's when you have frequent access on the same data that don't change often. By implementing cache on those data can reduce the load on your primary data store (database) as data cache searches faster than database searches.<br>
Example of common cache usage is in e-commerce for product catalog. Users frequently browse this catalog that includes name, price, desc of product. These product data doesn't change rapidly throughout the day, for this scenario, implementing caching system is the best strategy.
