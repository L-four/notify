// Code generated by mockery v2.14.0. DO NOT EDIT.

package discord

import (
	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// mockDiscordSession is an autogenerated mock type for the discordSession type
type mockDiscordSession struct {
	mock.Mock
}

// ChannelMessageSend provides a mock function with given fields: channelID, content
func (_m *mockDiscordSession) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	ret := _m.Called(channelID, content)

	var r0 *discordgo.Message
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Message); ok {
		r0 = rf(channelID, content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelID, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockDiscordSession interface {
	mock.TestingT
	Cleanup(func())
}

// newMockDiscordSession creates a new instance of mockDiscordSession. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockDiscordSession(t mockConstructorTestingTnewMockDiscordSession) *mockDiscordSession {
	mock := &mockDiscordSession{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
