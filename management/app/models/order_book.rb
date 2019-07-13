class OrderBook < ApplicationRecord
  enum status: %i[pending done cancel reject]

  belongs_to :user
  belongs_to :fund
end
