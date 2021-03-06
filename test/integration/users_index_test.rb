require 'test_helper'

class UsersIndexTest < ActionDispatch::IntegrationTest

    def setup
        @admin = users(:Michael)
        @non_admin = users(:Noel)
        @user = users(:Tsubaki)
    end

    test "index including pagination and delete links" do
        log_in_as(@admin)
        get users_path
        assert_template 'users/index'
        assert_select 'div.pagination'
        User.where(activated:true).paginate(page:1).each do |user|
            assert_select 'a[href=?]',user_path(user), text:user.name
            unless user == @admin
                assert_select 'a[href=?]',user_path(user), text:'delete'
            end
        end
        assert_select 'a[href=?]','/users?page=2',count:4
        assert_difference "User.count",-1 do
            delete user_path(@non_admin)
        end
    end

    test "index asa non-admin" do
        log_in_as(@non_admin)
        get users_path
        assert_select 'a',text:'delete',count:0
    end

end
