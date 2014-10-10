# Justcoin CLI app

CLI app to quickly check Justcoin.com markets

## Installation

Install the app on your system:

    $ go get http://github.com/bencolon/justcoin

## Usage

Enjoy `justcoin` executable

```bash
Commands:
  justcoin setup, s                   # Setup your justcoin.com api key

  justcoin markets, l                 # List all the justcoin.com markets
    options
      --BTCEUR '1'                    # BTC vs EUR trade
      --BTCLTC '1'                    # BTC vs LTC trade
      --BTCNOK '1'                    # BTC vs NOK trade
      --BTCSTR '1'                    # BTC vs STR trade
      --BTCUSD '1'                    # BTC vs USD trade
      --BTCXRP '1'                    # BTC vs XRP trade
      --trend  'yes'                  # Add trend colors to the market values

  justcoin wallet, w                  # Display your wallet amounts depending the options
    options:
      --BTC   'amount'                # BTC amount [at least one crypto amount is mandatory]
      --LTC   'amount'                # LTC amount
      --STR   'amount'                # STR amount
      --XRP   'amount'                # XRP amount
      --curr  'currency'              # Wallet curency : EUR (default) or USD or NOK) [mandatory]
      --tot   'yes'                   # Display one more line with the wallet total

  justcoin help, h                    # Shows a list of commands or help for one command
```

## Examples

```bash
Commands:
  justcoin markets
    BTCEUR = 274.23
    BTCLTC = 85.87
    BTCNOK = 2241.00
    BTCSTR = 142980.00
    BTCUSD = 344.63
    BTCXRP = 72000.72

  justcoin markets --BTCEUR 1 --BTCUSD 1
    BTCEUR = 274.23
    BTCUSD = 344.63

  justcoin wallet --BTC 12 --currency EUR
    12.00 BTC = 3290.76 EUR

  justcoin wallet --BTC 10 --LTC 50 --currency USD --tot yes
    10.00 BTC = 2742.30 EUR
    50.00 LTC = 159.67 EUR
    TOTAL = 2901.97 EUR
```

## Contributing

1. Fork it ( https://github.com/bencolon/justcoin/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request




