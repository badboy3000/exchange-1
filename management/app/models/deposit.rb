class Deposit < ApplicationRecord
  belongs_to :account
  belongs_to :currency
end
