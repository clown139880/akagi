class Caseship < ApplicationRecord
    belongs_to :father, class_name: 'Case'
    belongs_to :sub, class_name: 'Case'
end
