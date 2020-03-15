package createinvite

import (
    "github.com/jonas747/dcmd"
    "github.com/jonas747/discordgo"
    "github.com/jonas747/yagpdb/bot"
    "github.com/jonas747/yagpdb/commands"
    "github.com/jonas747/yagpdb/common"
)

var Command = &commands.YAGCommand{
    Cooldown:             30,
    CmdCategory:          commands.CategoryTool,
    HideFromCommandsPage: true,
    Name:                 "createinvite",
    Description:          "Creates an invite for the specified channel in the current server.",
    HideFromHelp:         false,
    RequiredArgs:         1,
    Arguments: []*dcmd.ArgDef{
        {Name: "channel", Type: dcmd.Int},
    },
    RunFunc: func(data *dcmd.Data) (interface{}, error) {
        if ok, err := bot.AdminOrPerm(0, data.Msg.Author.ID, data.CS.ID); err != nil {
            return "Failed checking perms", err
        } else if !ok {
            return "You need server admin perms to use this command", nil
        }

        invite, err := common.BotSession.ChannelInviteCreate(data.Args[0].Int64(), discordgo.Invite{
            MaxAge:  3600,
            MaxUses: 1,
        })

        if err != nil {
            return nil, err
        }

        return ("https://discord.gg/"+invite.Code), nil
    },
}
