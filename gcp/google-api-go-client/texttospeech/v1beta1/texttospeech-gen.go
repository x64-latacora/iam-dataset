// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package texttospeech provides access to the Cloud Text-to-Speech API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/texttospeech/apiv1 instead.
//
// For product documentation, see: https://cloud.google.com/text-to-speech/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/texttospeech/v1beta1"
//   ...
//   ctx := context.Background()
//   texttospeechService, err := texttospeech.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   texttospeechService, err := texttospeech.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   texttospeechService, err := texttospeech.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package texttospeech // import "google.golang.org/api/texttospeech/v1beta1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "texttospeech:v1beta1"
const apiName = "texttospeech"
const apiVersion = "v1beta1"
const basePath = "https://texttospeech.googleapis.com/"
const mtlsBasePath = "https://texttospeech.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Text = NewTextService(s)
	s.Voices = NewVoicesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Text *TextService

	Voices *VoicesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewTextService(s *Service) *TextService {
	rs := &TextService{s: s}
	return rs
}

type TextService struct {
	s *Service
}

func NewVoicesService(s *Service) *VoicesService {
	rs := &VoicesService{s: s}
	return rs
}

type VoicesService struct {
	s *Service
}

// AudioConfig: Description of audio data to be synthesized.
type AudioConfig struct {
	// AudioEncoding: Required. The format of the audio byte stream.
	//
	// Possible values:
	//   "AUDIO_ENCODING_UNSPECIFIED" - Not specified. Will return result
	// google.rpc.Code.INVALID_ARGUMENT.
	//   "LINEAR16" - Uncompressed 16-bit signed little-endian samples
	// (Linear PCM). Audio content returned as LINEAR16 also contains a WAV
	// header.
	//   "MP3" - MP3 audio at 32kbps.
	//   "MP3_64_KBPS" - MP3 at 64kbps.
	//   "OGG_OPUS" - Opus encoded audio wrapped in an ogg container. The
	// result will be a file which can be played natively on Android, and in
	// browsers (at least Chrome and Firefox). The quality of the encoding
	// is considerably higher than MP3 while using approximately the same
	// bitrate.
	//   "MULAW" - 8-bit samples that compand 14-bit audio samples using
	// G.711 PCMU/mu-law. Audio content returned as MULAW also contains a
	// WAV header.
	//   "ALAW" - 8-bit samples that compand 14-bit audio samples using
	// G.711 PCMU/A-law. Audio content returned as ALAW also contains a WAV
	// header.
	AudioEncoding string `json:"audioEncoding,omitempty"`

	// EffectsProfileId: Optional. Input only. An identifier which selects
	// 'audio effects' profiles that are applied on (post synthesized) text
	// to speech. Effects are applied on top of each other in the order they
	// are given. See audio profiles
	// (https://cloud.google.com/text-to-speech/docs/audio-profiles) for
	// current supported profile ids.
	EffectsProfileId []string `json:"effectsProfileId,omitempty"`

	// Pitch: Optional. Input only. Speaking pitch, in the range [-20.0,
	// 20.0]. 20 means increase 20 semitones from the original pitch. -20
	// means decrease 20 semitones from the original pitch.
	Pitch float64 `json:"pitch,omitempty"`

	// SampleRateHertz: Optional. The synthesis sample rate (in hertz) for
	// this audio. When this is specified in SynthesizeSpeechRequest, if
	// this is different from the voice's natural sample rate, then the
	// synthesizer will honor this request by converting to the desired
	// sample rate (which might result in worse audio quality), unless the
	// specified sample rate is not supported for the encoding chosen, in
	// which case it will fail the request and return
	// google.rpc.Code.INVALID_ARGUMENT.
	SampleRateHertz int64 `json:"sampleRateHertz,omitempty"`

	// SpeakingRate: Optional. Input only. Speaking rate/speed, in the range
	// [0.25, 4.0]. 1.0 is the normal native speed supported by the specific
	// voice. 2.0 is twice as fast, and 0.5 is half as fast. If unset(0.0),
	// defaults to the native 1.0 speed. Any other values < 0.25 or > 4.0
	// will return an error.
	SpeakingRate float64 `json:"speakingRate,omitempty"`

	// VolumeGainDb: Optional. Input only. Volume gain (in dB) of the normal
	// native volume supported by the specific voice, in the range [-96.0,
	// 16.0]. If unset, or set to a value of 0.0 (dB), will play at normal
	// native signal amplitude. A value of -6.0 (dB) will play at
	// approximately half the amplitude of the normal native signal
	// amplitude. A value of +6.0 (dB) will play at approximately twice the
	// amplitude of the normal native signal amplitude. Strongly recommend
	// not to exceed +10 (dB) as there's usually no effective increase in
	// loudness for any value greater than that.
	VolumeGainDb float64 `json:"volumeGainDb,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioEncoding") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioEncoding") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AudioConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AudioConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *AudioConfig) UnmarshalJSON(data []byte) error {
	type NoMethod AudioConfig
	var s1 struct {
		Pitch        gensupport.JSONFloat64 `json:"pitch"`
		SpeakingRate gensupport.JSONFloat64 `json:"speakingRate"`
		VolumeGainDb gensupport.JSONFloat64 `json:"volumeGainDb"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Pitch = float64(s1.Pitch)
	s.SpeakingRate = float64(s1.SpeakingRate)
	s.VolumeGainDb = float64(s1.VolumeGainDb)
	return nil
}

// CustomVoiceParams: Description of the custom voice to be synthesized.
type CustomVoiceParams struct {
	// Model: Required. The name of the AutoML model that synthesizes the
	// custom voice.
	Model string `json:"model,omitempty"`

	// ReportedUsage: Optional. The usage of the synthesized audio to be
	// reported.
	//
	// Possible values:
	//   "REPORTED_USAGE_UNSPECIFIED" - Request with reported usage
	// unspecified will be rejected.
	//   "REALTIME" - For scenarios where the synthesized audio is not
	// downloadable and can only be used once. For example, real-time
	// request in IVR system.
	//   "OFFLINE" - For scenarios where the synthesized audio is
	// downloadable and can be reused. For example, the synthesized audio is
	// downloaded, stored in customer service system and played repeatedly.
	ReportedUsage string `json:"reportedUsage,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Model") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Model") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomVoiceParams) MarshalJSON() ([]byte, error) {
	type NoMethod CustomVoiceParams
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListVoicesResponse: The message returned to the client by the
// `ListVoices` method.
type ListVoicesResponse struct {
	// Voices: The list of voices.
	Voices []*Voice `json:"voices,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Voices") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Voices") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListVoicesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListVoicesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SynthesisInput: Contains text input to be synthesized. Either `text`
// or `ssml` must be supplied. Supplying both or neither returns
// google.rpc.Code.INVALID_ARGUMENT. The input size is limited to 5000
// characters.
type SynthesisInput struct {
	// Ssml: The SSML document to be synthesized. The SSML document must be
	// valid and well-formed. Otherwise the RPC will fail and return
	// google.rpc.Code.INVALID_ARGUMENT. For more information, see SSML
	// (https://cloud.google.com/text-to-speech/docs/ssml).
	Ssml string `json:"ssml,omitempty"`

	// Text: The raw text to be synthesized.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Ssml") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Ssml") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SynthesisInput) MarshalJSON() ([]byte, error) {
	type NoMethod SynthesisInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SynthesizeSpeechRequest: The top-level message sent by the client for
// the `SynthesizeSpeech` method.
type SynthesizeSpeechRequest struct {
	// AudioConfig: Required. The configuration of the synthesized audio.
	AudioConfig *AudioConfig `json:"audioConfig,omitempty"`

	// EnableTimePointing: Whether and what timepoints are returned in the
	// response.
	//
	// Possible values:
	//   "TIMEPOINT_TYPE_UNSPECIFIED" - Not specified. No timepoint
	// information will be returned.
	//   "SSML_MARK" - Timepoint information of `` tags in SSML input will
	// be returned.
	EnableTimePointing []string `json:"enableTimePointing,omitempty"`

	// Input: Required. The Synthesizer requires either plain text or SSML
	// as input.
	Input *SynthesisInput `json:"input,omitempty"`

	// Voice: Required. The desired voice of the synthesized audio.
	Voice *VoiceSelectionParams `json:"voice,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SynthesizeSpeechRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SynthesizeSpeechRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SynthesizeSpeechResponse: The message returned to the client by the
// `SynthesizeSpeech` method.
type SynthesizeSpeechResponse struct {
	// AudioConfig: The audio metadata of `audio_content`.
	AudioConfig *AudioConfig `json:"audioConfig,omitempty"`

	// AudioContent: The audio data bytes encoded as specified in the
	// request, including the header for encodings that are wrapped in
	// containers (e.g. MP3, OGG_OPUS). For LINEAR16 audio, we include the
	// WAV header. Note: as with all bytes fields, protobuffers use a pure
	// binary representation, whereas JSON representations use base64.
	AudioContent string `json:"audioContent,omitempty"`

	// Timepoints: A link between a position in the original request input
	// and a corresponding time in the output audio. It's only supported via
	// `` of SSML input.
	Timepoints []*Timepoint `json:"timepoints,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AudioConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SynthesizeSpeechResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SynthesizeSpeechResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Timepoint: This contains a mapping between a certain point in the
// input text and a corresponding time in the output audio.
type Timepoint struct {
	// MarkName: Timepoint name as received from the client within `` tag.
	MarkName string `json:"markName,omitempty"`

	// TimeSeconds: Time offset in seconds from the start of the synthesized
	// audio.
	TimeSeconds float64 `json:"timeSeconds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MarkName") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MarkName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Timepoint) MarshalJSON() ([]byte, error) {
	type NoMethod Timepoint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Timepoint) UnmarshalJSON(data []byte) error {
	type NoMethod Timepoint
	var s1 struct {
		TimeSeconds gensupport.JSONFloat64 `json:"timeSeconds"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.TimeSeconds = float64(s1.TimeSeconds)
	return nil
}

// Voice: Description of a voice supported by the TTS service.
type Voice struct {
	// LanguageCodes: The languages that this voice supports, expressed as
	// BCP-47 (https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tags
	// (e.g. "en-US", "es-419", "cmn-tw").
	LanguageCodes []string `json:"languageCodes,omitempty"`

	// Name: The name of this voice. Each distinct voice has a unique name.
	Name string `json:"name,omitempty"`

	// NaturalSampleRateHertz: The natural sample rate (in hertz) for this
	// voice.
	NaturalSampleRateHertz int64 `json:"naturalSampleRateHertz,omitempty"`

	// SsmlGender: The gender of this voice.
	//
	// Possible values:
	//   "SSML_VOICE_GENDER_UNSPECIFIED" - An unspecified gender. In
	// VoiceSelectionParams, this means that the client doesn't care which
	// gender the selected voice will have. In the Voice field of
	// ListVoicesResponse, this may mean that the voice doesn't fit any of
	// the other categories in this enum, or that the gender of the voice
	// isn't known.
	//   "MALE" - A male voice.
	//   "FEMALE" - A female voice.
	//   "NEUTRAL" - A gender-neutral voice. This voice is not yet
	// supported.
	SsmlGender string `json:"ssmlGender,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCodes") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCodes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Voice) MarshalJSON() ([]byte, error) {
	type NoMethod Voice
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VoiceSelectionParams: Description of which voice to use for a
// synthesis request.
type VoiceSelectionParams struct {
	// CustomVoice: The configuration for a custom voice. If
	// [CustomVoiceParams.model] is set, the service will choose the custom
	// voice matching the specified configuration.
	CustomVoice *CustomVoiceParams `json:"customVoice,omitempty"`

	// LanguageCode: Required. The language (and potentially also the
	// region) of the voice expressed as a BCP-47
	// (https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag, e.g.
	// "en-US". This should not include a script tag (e.g. use "cmn-cn"
	// rather than "cmn-Hant-cn"), because the script will be inferred from
	// the input provided in the SynthesisInput. The TTS service will use
	// this parameter to help choose an appropriate voice. Note that the TTS
	// service may choose a voice with a slightly different language code
	// than the one selected; it may substitute a different region (e.g.
	// using en-US rather than en-CA if there isn't a Canadian voice
	// available), or even a different language, e.g. using "nb" (Norwegian
	// Bokmal) instead of "no" (Norwegian)".
	LanguageCode string `json:"languageCode,omitempty"`

	// Name: The name of the voice. If not set, the service will choose a
	// voice based on the other parameters such as language_code and gender.
	Name string `json:"name,omitempty"`

	// SsmlGender: The preferred gender of the voice. If not set, the
	// service will choose a voice based on the other parameters such as
	// language_code and name. Note that this is only a preference, not
	// requirement; if a voice of the appropriate gender is not available,
	// the synthesizer should substitute a voice with a different gender
	// rather than failing the request.
	//
	// Possible values:
	//   "SSML_VOICE_GENDER_UNSPECIFIED" - An unspecified gender. In
	// VoiceSelectionParams, this means that the client doesn't care which
	// gender the selected voice will have. In the Voice field of
	// ListVoicesResponse, this may mean that the voice doesn't fit any of
	// the other categories in this enum, or that the gender of the voice
	// isn't known.
	//   "MALE" - A male voice.
	//   "FEMALE" - A female voice.
	//   "NEUTRAL" - A gender-neutral voice. This voice is not yet
	// supported.
	SsmlGender string `json:"ssmlGender,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomVoice") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomVoice") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VoiceSelectionParams) MarshalJSON() ([]byte, error) {
	type NoMethod VoiceSelectionParams
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "texttospeech.text.synthesize":

type TextSynthesizeCall struct {
	s                       *Service
	synthesizespeechrequest *SynthesizeSpeechRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// Synthesize: Synthesizes speech synchronously: receive results after
// all text input has been processed.
func (r *TextService) Synthesize(synthesizespeechrequest *SynthesizeSpeechRequest) *TextSynthesizeCall {
	c := &TextSynthesizeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.synthesizespeechrequest = synthesizespeechrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TextSynthesizeCall) Fields(s ...googleapi.Field) *TextSynthesizeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TextSynthesizeCall) Context(ctx context.Context) *TextSynthesizeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TextSynthesizeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TextSynthesizeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211120")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.synthesizespeechrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/text:synthesize")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "texttospeech.text.synthesize" call.
// Exactly one of *SynthesizeSpeechResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *SynthesizeSpeechResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TextSynthesizeCall) Do(opts ...googleapi.CallOption) (*SynthesizeSpeechResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SynthesizeSpeechResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Synthesizes speech synchronously: receive results after all text input has been processed.",
	//   "flatPath": "v1beta1/text:synthesize",
	//   "httpMethod": "POST",
	//   "id": "texttospeech.text.synthesize",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/text:synthesize",
	//   "request": {
	//     "$ref": "SynthesizeSpeechRequest"
	//   },
	//   "response": {
	//     "$ref": "SynthesizeSpeechResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "texttospeech.voices.list":

type VoicesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns a list of Voice supported for synthesis.
func (r *VoicesService) List() *VoicesListCall {
	c := &VoicesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// LanguageCode sets the optional parameter "languageCode": Recommended.
// BCP-47 (https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag.
// If not specified, the API will return all supported voices. If
// specified, the ListVoices call will only return voices that can be
// used to synthesize this language_code. E.g. when specifying
// "en-NZ", you will get supported "en-NZ" voices; when specifying
// "no", you will get supported "no-\*" (Norwegian) and "nb-\*"
// (Norwegian Bokmal) voices; specifying "zh" will also get supported
// "cmn-\*" voices; specifying "zh-hk" will also get supported
// "yue-hk" voices.
func (c *VoicesListCall) LanguageCode(languageCode string) *VoicesListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VoicesListCall) Fields(s ...googleapi.Field) *VoicesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *VoicesListCall) IfNoneMatch(entityTag string) *VoicesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VoicesListCall) Context(ctx context.Context) *VoicesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VoicesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VoicesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211120")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/voices")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "texttospeech.voices.list" call.
// Exactly one of *ListVoicesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListVoicesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *VoicesListCall) Do(opts ...googleapi.CallOption) (*ListVoicesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListVoicesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of Voice supported for synthesis.",
	//   "flatPath": "v1beta1/voices",
	//   "httpMethod": "GET",
	//   "id": "texttospeech.voices.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. Recommended. [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. If not specified, the API will return all supported voices. If specified, the ListVoices call will only return voices that can be used to synthesize this language_code. E.g. when specifying `\"en-NZ\"`, you will get supported `\"en-NZ\"` voices; when specifying `\"no\"`, you will get supported `\"no-\\*\"` (Norwegian) and `\"nb-\\*\"` (Norwegian Bokmal) voices; specifying `\"zh\"` will also get supported `\"cmn-\\*\"` voices; specifying `\"zh-hk\"` will also get supported `\"yue-hk\"` voices.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/voices",
	//   "response": {
	//     "$ref": "ListVoicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
