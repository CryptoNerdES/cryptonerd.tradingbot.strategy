# cn.example.cryptobot.dca

TDD and Go to create a CryptoTradingBot for Binance using DCA

The idea behind this strategy is not only doing DCA (Dolar cost average) is selling and buying when we are winning a trade.

One example:
```
Day 1:We buy 50$ in BTC (0,00128 BTC) at 7:00 AM
Day 2: BTC goes up 3% at 7:00 AM, we sell the position. Nice trade
Day 3: we buy again 50$ at 7:00 AM
Day 4: BTC goes down 1% at 7:00 AM, we buy again 50$
Day 5: BTC goes down 2% at 7:00 AM, we buy again 50$
Day 6: BTC goes up 4% at 7:00 AM, we sell the position. Nice trade
```

The idea is to sell BTC goes up of our entry price, and buy while we are losing, reducing the entry price of our trade.

This Strategy is only for BTC, this is because the other cryptocurrencies are not enough stable. You can try with ETH but is not recommended to do with other altcoins.