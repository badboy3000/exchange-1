# frozen_string_literal: true

require_relative 'boot'

require 'rails'
# Pick the frameworks you want:
require 'active_model/railtie'
require 'active_job/railtie'
require 'active_record/railtie'
# require "active_storage/engine"
require 'action_controller/railtie'
require 'action_mailer/railtie'
# require "action_mailbox/engine"
# require "action_text/engine"
require 'action_view/railtie'
# require "action_cable/engine"
require 'sprockets/railtie'
# require "rails/test_unit/railtie"

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module Management
  class Application < Rails::Application
    # Initialize configuration defaults for originally generated Rails version.
    config.load_defaults 6.0

    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration can go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded after loading
    # the framework and any gems in your application.
    config.autoload_paths << Rails.root.join('lib')
    config.eager_load_paths << Rails.root.join('lib')

    config.generators do |g|
      g.orm :active_record
      g.template_engine nil
      # g.test_framework :rspec,
      #   view_specs: false,
      #   helper_specs: false,
      #   routing_specs: false,
      #   controller_specs: false,
      #   model_specs: false,
      #   request_specs: true,
      #   mailer_specs: false,
      #   job_spec: false,
      #   system_specs: false
      # g.fixture_replacement :factory_bot, dir: 'spec/factories'
      g.test_framework nil
      g.controller_specs false
      g.request_specs false
      g.job_specs false
      g.helper_specs false
      g.feature_specs false
      g.mailer_specs false
      g.model_specs false
      g.observer_specs false
      g.routing_specs false
      g.view_specs false

      g.factory_bot false
      g.integration_tool :rspec
      g.stylesheets false
      g.javascripts false
      g.jbuilder false
      g.helper false
    end

    config.i18n.available_locales = ['zh-CN', :en]
    config.i18n.default_locale = 'zh-CN'

    config.active_record.default_timezone = :local
    config.time_zone = 'Beijing'
    config.encoding = 'utf-8'
  end
end
