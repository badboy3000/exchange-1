# frozen_string_literal: true

huobi = Huobi.new
huobi.currencys['data'].each do |s|
  Currency.create!(
    symbol: s,
    deposit_fee: 0.0001,
    withdraw_fee: 0.0001
  )
end

Huobi.new.symbols['data'].each do |s|
  base_currency = s['base-currency']
  quote_currency = s['quote-currency']
  symbol = "#{base_currency}_#{quote_currency}"
  Fund.create!(
    name: symbol,
    symbol: symbol,
    left_currency_id: Currency.where(symbol: base_currency).last&.id,
    right_currency_id: Currency.where(symbol: quote_currency).last&.id,
    limit_rate: 0.0001,
    market_rate: 0.0002
  )
end

btc_usdt = Fund.where(symbol: 'btc_usdt').last
Huobi.new.history_trade('btcusdt', 2000)['data'].each do |row|
  row['data'].each do |t|
    Order.create!(
      bid_user_id: 1,
      ask_user_id: 1,
      symbol: btc_usdt&.symbol,
      fund_id: btc_usdt&.id,
      bid_order_book_id: 1,
      ask_order_book_id: 1,
      volume: t['amount'],
      price: t['price'],
      created_at: DateTime.strptime((t['ts']).to_s, '%Q')
    )
  end
end

user = User.new(
  name: 'yang',
  password: '111111',
  password_confirmation: '111111',
  email: 'sysuyangkang@gmail.com',
  role: 'customer',
  address: '北京朝阳'
)
user.save!

Currency.all.each do |c|
  account = Account.create!(
    user_id: user.id,
    currency_id: c.id,
    balance: 0,
    locked: 0
  )

  Deposit.create!(
    account_id: account.id,
    currency_id: c.id,
    amount: 100,
    fee: 0
  )
end

Account.all.each do |a|
  a.update!(balance: Deposit.where(account_id: a.id, currency_id: a.currency_id).sum(:amount))
end
