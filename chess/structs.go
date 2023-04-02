package chess

// BearerResponse - a
type BearerResponse struct {
	Status string `json:"status"`
	Data   struct {
		UUID            string      `json:"uuid"`
		IsGuest         bool        `json:"is_guest"`
		Email           interface{} `json:"email"`
		ObfuscatedEmail interface{} `json:"obfuscated_email"`
		Cohort          interface{} `json:"cohort"`
		HasLcPriority   bool        `json:"has_lc_priority"`
		IsStreamer      bool        `json:"is_streamer"`
		Oauth           struct {
			TokenType    string `json:"token_type"`
			ExpiresIn    int    `json:"expires_in"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"oauth"`
		LoginToken    string      `json:"login_token"`
		PremiumStatus int         `json:"premium_status"`
		ID            int         `json:"id"`
		CountryID     int         `json:"country_id"`
		AvatarURL     string      `json:"avatar_url"`
		LastLoginDate int         `json:"last_login_date"`
		SessionID     string      `json:"session_id"`
		Location      interface{} `json:"location"`
		Username      string      `json:"username"`
		ChessTitle    interface{} `json:"chess_title"`
		MemberSince   int         `json:"member_since"`
		FlairCode     string      `json:"flair_code"`
		ShowAds       bool        `json:"show_ads"`
		Config        struct {
			Live struct {
				Connections []string `json:"connections"`
			} `json:"live"`
			Integration struct {
				Amplitude struct {
					Key string `json:"key"`
				} `json:"amplitude"`
				Iterable struct {
					Key string `json:"key"`
				} `json:"iterable"`
			} `json:"integration"`
			Features       []string `json:"features"`
			FeaturesSignal struct {
				PleaseRateDialog                                 interface{} `json:"PleaseRateDialog"`
				PubSub                                           interface{} `json:"pub_sub"`
				RcnGuestPlay                                     interface{} `json:"rcn_guest_play"`
				RcnUnratedPlay                                   interface{} `json:"rcn_unrated_play"`
				ConnectedBoards                                  interface{} `json:"connected_boards"`
				InsightsWebview                                  interface{} `json:"insights_webview"`
				LeaguesWebview                                   interface{} `json:"leagues_webview"`
				GameReview                                       interface{} `json:"game_review"`
				RcnGameChat                                      interface{} `json:"rcn_game_chat"`
				LessonsConsumeQuotaWhenUserCompletesOneChallenge interface{} `json:"lessons_consume_quota_when_user_completes_one_challenge"`
				RussianFlagNowar                                 interface{} `json:"russian_flag_nowar"`
				RatedPuzzlesRewardedVideo                        interface{} `json:"rated_puzzles_rewarded_video"`
				PremiumTrialOnboarding                           interface{} `json:"premium_trial_onboarding"`
				AdsPerformanceTrackingKillswitch                 interface{} `json:"ads_performance_tracking_killswitch"`
				ChessTvHomeRedesign                              interface{} `json:"chess_tv_home_redesign"`
				CheatingDetection                                interface{} `json:"cheating_detection"`
				PricingChange2022                                interface{} `json:"pricing_change_2022"`
				NewOnboardingSeptember2022                       interface{} `json:"new_onboarding_september2022"`
				CheckRequiredGameMetrics                         interface{} `json:"check_required_game_metrics"`
				DisconnectedGameMetrics                          interface{} `json:"disconnected_game_metrics"`
				RefreshAuthTokenInWebviews                       interface{} `json:"refresh_auth_token_in_webviews"`
				DatadogExtraLogsExperiment                       interface{} `json:"datadog_extra_logs_experiment"`
				SyncFriendRequests                               interface{} `json:"sync_friend_requests"`
				OpeningStatsKillswitch                           interface{} `json:"opening_stats_killswitch"`
				RcnRatedGames                                    interface{} `json:"rcn_rated_games"`
				OnboardingTrialExperiment                        interface{} `json:"onboarding_trial_experiment"`
				DailyOutgoingChallengeHomeItem                   interface{} `json:"daily_outgoing_challenge_home_item"`
				HomeGoPremiumExperiment                          interface{} `json:"home_go_premium_experiment"`
			} `json:"features_signal"`
		} `json:"config"`
		IsFairPlayAgreed bool `json:"is_fair_play_agreed"`
	} `json:"data"`
}
