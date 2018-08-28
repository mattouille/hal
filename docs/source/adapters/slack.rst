=====
Slack
=====

Setup
~~~~~

Currently Hal will listen in on all public channels, or a custom list of channels if ``HAL_SLACK_CHANNELS`` is declared. Private groups
require the IRC gateway to work around a current limitation of the Slack
API. See `Using IRC Gateway`_. The RTM method is the best method to use.
Some advanced features like attachment uploading are not supported at this time.

Usage
~~~~~

.. code:: go

    // blank import to register adapter
    import _ "github.com/mattouille/hal/adapter/slack"

Configuration
~~~~~~~~~~~~~

::

    HAL_ADAPTER=slack               # The adapter
    HAL_SLACK_TOKEN=blah            # Your integration token
                                    # Default: none (required)
    HAL_SLACK_TEAM=acmeinc          # Your Slack subdomain (<team>.slack.com)
                                    # Default: none
    HAL_SLACK_BOTNAME=HAL           # The username Hal will send replies as
                                    # Default: HAL_NAME
    HAL_SLACK_CHANNELS=""           # Channels to whitelist or blakcklist
                                    # See HAL_SLACK_CHANNEL_MODE
    HAL_SLACK_CHANNEL_MODE=""       #
    HAL_SLACK_LINK_NAMES=""         # not yet implemented

``HAL_SLACK_CHANNEL_MODE``
^^^^^^^^^^^^^^^^^^^^^^^^

Specify how to treat the list of channels in ``HAL_SLACK_CHANNELS```. Disabled if ``HAL_SLACK_CHANNELS`` is empty
or ``HAL_SLACK_CHANNEL_MODE`` is ``none``.

- **Options:** whitelist, blacklist, none
- **Default:** none
- **Required:** false
- **Example:** ``HAL_SLACK_CHANNEL_MODE=whitelist``
-

Using IRC Gateway
^^^^^^^^^^^^^^^^^

The default integration only works with public chats. If you want hal to
listen in on private chats, you must utilize the IRC gateway. You'll
need a real user for hal, so be mindful of the username you choose for
it and make sure you configure your bot to use that name so it can login
to the IRC gateway. When enabled, hal will only use the IRC gateway to
listen for messages. Hal can be configured to either respond using the
API or the IRC gateway.

1. Enable the IRC gateway in `the admin settings
   interface <https://revily.slack.com/admin/settings>`__

   -  Choose "Enable IRC gateway (SSL only)". You don't want your
      private messages sent unencrypted.

2. `Register <https://my.slack.com/signup>`__ a new user
3. Sign in as this new user
4. Capture your new `IRC credentials <https://my.slack.com/account/gateways>`__
5. Set the following environment variables

::

    HAL_SLACK_IRC_ENABLED           # Enable the Slack IRC listener
                                    # Default: 0
                                    # Options: 0, 1  ; 0 is disabled, 1 is enabled
    HAL_SLACK_IRC_PASSWORD          # The IRC gateway password
                                    # Default: none (required)
    HAL_SLACK_RESPONSE_METHOD       # The method by which hal will respond to a message.
                                    # The irc option requires that the IRC gateway be configured
                                    # Default: rtm
                                    # Options: rtm, irc

For more information, please see the following link: \* `Connecting to
Slack over IRC and
XMPP <https://slack.zendesk.com/hc/en-us/articles/201727913-Connecting-to-Slack-over-IRC-and-XMPP>`__
