
%w[BTC ETH LTC XRP BCH EOS DASH TRX ONT IOST IOTA USD USDT].each do |s|
  Currency.create!(
    symbol: s,
    deposit_fee: 0.0001,
    withdraw_fee: 0.0001,
  )
end

usd_currency = Currency.where(symbol: 'USD').first
Currency.first(11).each do |c|
  Fund.create!(
    name: "#{c.symbol}_USD",
    symbol: "#{c.symbol}_USD",
    left_currency_id: c.id,
    right_currency_id: usd_currency.id,
    limit_rate: 0.0001,
    market_rate: 0.0002
  )
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
