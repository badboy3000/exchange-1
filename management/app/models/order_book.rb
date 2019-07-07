class OrderBook < ApplicationRecord
  belongs_to :user
  belongs_to :fund
end
