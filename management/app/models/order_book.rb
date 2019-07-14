class OrderBook < ApplicationRecord
  enum status: %i[pending done cancel]

  belongs_to :user
  belongs_to :fund
end
