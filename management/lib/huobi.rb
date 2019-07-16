# frozen_string_literal: true

require 'httparty'
require 'json'
require 'open-uri'
require 'rack'
require 'digest/md5'
require 'base64'

class Huobi
  def initialize
    @uri = URI.parse 'https://api.huobi.pro/'
    @header = {
      'Content-Type' => 'application/json',
      'Accept' => 'application/json',
      'Accept-Language' => 'zh-CN',
      'User-Agent' => 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36'
    }
  end

  ### https://huobiapi.github.io/docs/spot/v1/cn/#41654e0268
  ## 获取所有交易对
  def symbols
    path = '/v1/common/symbols'
    request_method = 'GET'
    params = {}
    util(path, params, request_method)
  end

  ## 获取所有币种
  def currencys
    path = '/v1/common/currencys'
    request_method = 'GET'
    params = {}
    util(path, params, request_method)
  end

  ## 获取当前系统时间 TODO

  ### https://huobiapi.github.io/docs/spot/v1/cn/#d9d514d202
  ## K 线数据（蜡烛图）
  def history_kline(symbol, period, size = 150)
    path = '/market/history/kline'
    request_method = 'GET'
    params = { 'symbol' => symbol, 'period' => period, 'size' => size }
    util(path, params, request_method)
  end

  ## 聚合行情（Ticker）
  def merged(symbol)
    path = '/market/detail/merged'
    request_method = 'GET'
    params = { 'symbol' => symbol }
    util(path, params, request_method)
  end

  ## 所有交易对的最新 Tickers TODO

  ## 市场深度数据
  def depth(symbol, type = 'step0')
    path = '/market/depth'
    request_method = 'GET'
    params = { 'symbol' => symbol, 'type' => type }
    util(path, params, request_method)
  end

  ## 最近市场成交记录
  def market_trade(symbol)
    path = '/market/depth'
    request_method = 'GET'
    params = { 'symbol' => symbol }
    util(path, params, request_method)
  end

  ## 获得近期交易记录
  def history_trade(symbol, size = 1)
    path = '/market/history/trade'
    request_method = 'GET'
    params = { 'symbol' => symbol, 'size' => size }
    util(path, params, request_method)
  end

  ## 最近24小时行情数据
  def market_detail(symbol)
    path = '/market/detail'
    request_method = 'GET'
    params = { 'symbol' => symbol }
    util(path, params, request_method)
  end

  private

  def util(path, params, request_method)
    h = {
      'Timestamp' => Time.now.getutc.strftime('%Y-%m-%dT%H:%M:%S')
    }
    h = h.merge(params) if request_method == 'GET'
    data = "#{request_method}\napi.huobi.pro\n#{path}\n#{Rack::Utils.build_query(hash_sort(h))}"
    url = "https://api.huobi.pro#{path}?#{Rack::Utils.build_query(h)}"
    http = Net::HTTP.new(@uri.host, @uri.port)
    http.use_ssl = true
    begin
      JSON.parse http.send_request(request_method, url, JSON.dump(params), @header).body
    rescue Exception => e
      { 'message' => 'error', 'request_error' => e.message }
    end
  end

  def hash_sort(ha)
    Hash[ha.sort_by { |key, _val| key }]
  end
end
