// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/cozely/internal"
)

//------------------------------------------------------------------------------

// A keyLabel designate a key by its label in the current layout of the keyboard.
// For printable characters, the value is the rune that would be generated by
// pressing the key without any modifiers.
type keyLabel = internal.KeyLabel

// keyLabel constants
const (
	keylabelUnknown keyLabel = 0

	keylabelReturn     keyLabel = '\r'
	keylabelEscape     keyLabel = '\033' //FIXME
	keylabelBackspace  keyLabel = '\b'
	keylabelTab        keyLabel = '\t'
	keylabelSpace      keyLabel = ' '
	keylabelExclaim    keyLabel = '!'
	keylabelQuoteDbl   keyLabel = '"'
	keylabelHash       keyLabel = '#'
	keylabelPercent    keyLabel = '%'
	keylabelDollar     keyLabel = '$'
	keylabelAmpersand  keyLabel = '&'
	keylabelQuote      keyLabel = '\''
	keylabelLeftParen  keyLabel = '('
	keylabelRightParen keyLabel = ')'
	keylabelAsterisk   keyLabel = '*'
	keylabelPlus       keyLabel = '+'
	keylabelComma      keyLabel = ','
	keylabelMinus      keyLabel = '-'
	keylabelPeriod     keyLabel = '.'
	keylabelSlash      keyLabel = '/'
	keylabel0          keyLabel = '0'
	keylabel1          keyLabel = '1'
	keylabel2          keyLabel = '2'
	keylabel3          keyLabel = '3'
	keylabel4          keyLabel = '4'
	keylabel5          keyLabel = '5'
	keylabel6          keyLabel = '6'
	keylabel7          keyLabel = '7'
	keylabel8          keyLabel = '8'
	keylabel9          keyLabel = '9'
	keylabelColon      keyLabel = ':'
	keylabelSemicolon  keyLabel = ';'
	keylabelLess       keyLabel = '<'
	keylabelEquals     keyLabel = '='
	keylabelGreater    keyLabel = '>'
	keylabelQuestion   keyLabel = '?'
	keylabelAt         keyLabel = '@'

	// Skip uppercase letters

	keylabelLeftBracket  keyLabel = '['
	keylabelBackSlash    keyLabel = '\\'
	keylabelRightBracket keyLabel = ']'
	keylabelCaret        keyLabel = '^'
	keylabelUnderscore   keyLabel = '_'
	keylabelBackQuote    keyLabel = '`'
	keylabelA            keyLabel = 'a'
	keylabelB            keyLabel = 'b'
	keylabelC            keyLabel = 'c'
	keylabelD            keyLabel = 'd'
	keylabelE            keyLabel = 'e'
	keylabelF            keyLabel = 'f'
	keylabelG            keyLabel = 'g'
	keylabelH            keyLabel = 'h'
	keylabelI            keyLabel = 'i'
	keylabelJ            keyLabel = 'j'
	keylabelK            keyLabel = 'k'
	keylabelL            keyLabel = 'l'
	keylabelM            keyLabel = 'm'
	keylabelN            keyLabel = 'n'
	keylabelO            keyLabel = 'o'
	keylabelP            keyLabel = 'p'
	keylabelQ            keyLabel = 'q'
	keylabelR            keyLabel = 'r'
	keylabelS            keyLabel = 's'
	keylabelT            keyLabel = 't'
	keylabelU            keyLabel = 'u'
	keylabelV            keyLabel = 'v'
	keylabelW            keyLabel = 'w'
	keylabelX            keyLabel = 'x'
	keylabelY            keyLabel = 'y'
	keylabelZ            keyLabel = 'z'

	keylabelCapsLock keyLabel = unprintable | keyLabel(keyCapsLock)

	keylabelF1  keyLabel = unprintable | keyLabel(keyF1)
	keylabelF2  keyLabel = unprintable | keyLabel(keyF2)
	keylabelF3  keyLabel = unprintable | keyLabel(keyF3)
	keylabelF4  keyLabel = unprintable | keyLabel(keyF4)
	keylabelF5  keyLabel = unprintable | keyLabel(keyF5)
	keylabelF6  keyLabel = unprintable | keyLabel(keyF6)
	keylabelF7  keyLabel = unprintable | keyLabel(keyF7)
	keylabelF8  keyLabel = unprintable | keyLabel(keyF8)
	keylabelF9  keyLabel = unprintable | keyLabel(keyF9)
	keylabelF10 keyLabel = unprintable | keyLabel(keyF10)
	keylabelF11 keyLabel = unprintable | keyLabel(keyF11)
	keylabelF12 keyLabel = unprintable | keyLabel(keyF12)

	keylabelPrintScreen keyLabel = unprintable | keyLabel(keyPrintScreen)
	keylabelScrollLock  keyLabel = unprintable | keyLabel(keyScrollLock)
	keylabelPause       keyLabel = unprintable | keyLabel(keyPause)
	keylabelInsert      keyLabel = unprintable | keyLabel(keyInsert)
	keylabelHome        keyLabel = unprintable | keyLabel(keyHome)
	keylabelPageUp      keyLabel = unprintable | keyLabel(keyPageUp)
	keylabelDelete      keyLabel = '\177' //FIXME
	keylabelEnd         keyLabel = unprintable | keyLabel(keyEnd)
	keylabelPageDown    keyLabel = unprintable | keyLabel(keyPageDown)
	keylabelRight       keyLabel = unprintable | keyLabel(keyRight)
	keylabelLeft        keyLabel = unprintable | keyLabel(keyLeft)
	keylabelDown        keyLabel = unprintable | keyLabel(keyDown)
	keylabelUp          keyLabel = unprintable | keyLabel(keyUp)

	keylabelNumLockClear keyLabel = unprintable | keyLabel(keyNumLockClear)
	keylabelKPDivide     keyLabel = unprintable | keyLabel(keyKPDivide)
	keylabelKPMultiply   keyLabel = unprintable | keyLabel(keyKPMultiply)
	keylabelKPMinus      keyLabel = unprintable | keyLabel(keyKPMinus)
	keylabelKPPlus       keyLabel = unprintable | keyLabel(keyKPPlus)
	keylabelKPEnter      keyLabel = unprintable | keyLabel(keyKPEnter)
	keylabelKP1          keyLabel = unprintable | keyLabel(keyKP1)
	keylabelKP2          keyLabel = unprintable | keyLabel(keyKP2)
	keylabelKP3          keyLabel = unprintable | keyLabel(keyKP3)
	keylabelKP4          keyLabel = unprintable | keyLabel(keyKP4)
	keylabelKP5          keyLabel = unprintable | keyLabel(keyKP5)
	keylabelKP6          keyLabel = unprintable | keyLabel(keyKP6)
	keylabelKP7          keyLabel = unprintable | keyLabel(keyKP7)
	keylabelKP8          keyLabel = unprintable | keyLabel(keyKP8)
	keylabelKP9          keyLabel = unprintable | keyLabel(keyKP9)
	keylabelKP0          keyLabel = unprintable | keyLabel(keyKP0)
	keylabelKPPeriod     keyLabel = unprintable | keyLabel(keyKPPeriod)

	keylabelApplication   keyLabel = unprintable | keyLabel(keyApplication)
	keylabelPower         keyLabel = unprintable | keyLabel(keyPower)
	keylabelKPEquals      keyLabel = unprintable | keyLabel(keyKPEquals)
	keylabelF13           keyLabel = unprintable | keyLabel(keyF13)
	keylabelF14           keyLabel = unprintable | keyLabel(keyF14)
	keylabelF15           keyLabel = unprintable | keyLabel(keyF15)
	keylabelF16           keyLabel = unprintable | keyLabel(keyF16)
	keylabelF17           keyLabel = unprintable | keyLabel(keyF17)
	keylabelF18           keyLabel = unprintable | keyLabel(keyF18)
	keylabelF19           keyLabel = unprintable | keyLabel(keyF19)
	keylabelF20           keyLabel = unprintable | keyLabel(keyF20)
	keylabelF21           keyLabel = unprintable | keyLabel(keyF21)
	keylabelF22           keyLabel = unprintable | keyLabel(keyF22)
	keylabelF23           keyLabel = unprintable | keyLabel(keyF23)
	keylabelF24           keyLabel = unprintable | keyLabel(keyF24)
	keylabelExecute       keyLabel = unprintable | keyLabel(keyExecute)
	keylabelHelp          keyLabel = unprintable | keyLabel(keyHelp)
	keylabelMenu          keyLabel = unprintable | keyLabel(keyMenu)
	keylabelSelect        keyLabel = unprintable | keyLabel(keySelect)
	keylabelStop          keyLabel = unprintable | keyLabel(keyStop)
	keylabelAgain         keyLabel = unprintable | keyLabel(keyAgain)
	keylabelUndo          keyLabel = unprintable | keyLabel(keyUndo)
	keylabelCut           keyLabel = unprintable | keyLabel(keyCut)
	keylabelCopy          keyLabel = unprintable | keyLabel(keyCopy)
	keylabelPaste         keyLabel = unprintable | keyLabel(keyPaste)
	keylabelFind          keyLabel = unprintable | keyLabel(keyFind)
	keylabelMute          keyLabel = unprintable | keyLabel(keyMute)
	keylabelVolumeUp      keyLabel = unprintable | keyLabel(keyVolumeUp)
	keylabelVolumeDown    keyLabel = unprintable | keyLabel(keyVolumeDown)
	keylabelKPComma       keyLabel = unprintable | keyLabel(keyKPComma)
	keylabelKPEqualsAS400 keyLabel = unprintable | keyLabel(keyKPEqualsAS400)

	keylabelAltErase   keyLabel = unprintable | keyLabel(keyAltErase)
	keylabelSysReq     keyLabel = unprintable | keyLabel(keySysReq)
	keylabelCancel     keyLabel = unprintable | keyLabel(keyCancel)
	keylabelClear      keyLabel = unprintable | keyLabel(keyClear)
	keylabelPrior      keyLabel = unprintable | keyLabel(keyPrior)
	keylabelReturn2    keyLabel = unprintable | keyLabel(keyReturn2)
	keylabelSeparator  keyLabel = unprintable | keyLabel(keySeparator)
	keylabelOut        keyLabel = unprintable | keyLabel(keyOut)
	keylabelOper       keyLabel = unprintable | keyLabel(keyOper)
	keylabelClearAgain keyLabel = unprintable | keyLabel(keyClearAgain)
	keylabelCrSel      keyLabel = unprintable | keyLabel(keyCrSel)
	keylabelExSel      keyLabel = unprintable | keyLabel(keyExSel)

	keylabelKP00               keyLabel = unprintable | keyLabel(keyKP00)
	keylabelKP000              keyLabel = unprintable | keyLabel(keyKP000)
	keylabelThousandsSeparator keyLabel = unprintable | keyLabel(keyThousandsSeparator)
	keylabelDecimalSeparator   keyLabel = unprintable | keyLabel(keyDecimalSeparator)
	keylabelCurrencyUnit       keyLabel = unprintable | keyLabel(keyCurrencyUnit)
	keylabelCurrencySubUnit    keyLabel = unprintable | keyLabel(keyCurrencySubUnit)
	keylabelKPLeftParen        keyLabel = unprintable | keyLabel(keyKPLeftParen)
	keylabelKPRightParen       keyLabel = unprintable | keyLabel(keyKPRightParen)
	keylabelKPLeftBrace        keyLabel = unprintable | keyLabel(keyKPLeftBrace)
	keylabelKPRightBrace       keyLabel = unprintable | keyLabel(keyKPRightBrace)
	keylabelKPTab              keyLabel = unprintable | keyLabel(keyKPTab)
	keylabelKPBackspace        keyLabel = unprintable | keyLabel(keyKPBackspace)
	keylabelKPA                keyLabel = unprintable | keyLabel(keyKPA)
	keylabelKPB                keyLabel = unprintable | keyLabel(keyKPB)
	keylabelKPC                keyLabel = unprintable | keyLabel(keyKPC)
	keylabelKPD                keyLabel = unprintable | keyLabel(keyKPD)
	keylabelKPE                keyLabel = unprintable | keyLabel(keyKPE)
	keylabelKPF                keyLabel = unprintable | keyLabel(keyKPF)
	keylabelKPXor              keyLabel = unprintable | keyLabel(keyKPXor)
	keylabelKPPower            keyLabel = unprintable | keyLabel(keyKPPower)
	keylabelKPPercent          keyLabel = unprintable | keyLabel(keyKPPercent)
	keylabelKPLess             keyLabel = unprintable | keyLabel(keyKPLess)
	keylabelKPGreater          keyLabel = unprintable | keyLabel(keyKPGreater)
	keylabelKPAmpersand        keyLabel = unprintable | keyLabel(keyKPAmpersand)
	keylabelKPDblAmpersand     keyLabel = unprintable | keyLabel(keyKPDblAmpersand)
	keylabelKPVerticalBar      keyLabel = unprintable | keyLabel(keyKPVerticalBar)
	keylabelKPDblVerticalBar   keyLabel = unprintable | keyLabel(keyKPDblVerticalBar)
	keylabelKPColon            keyLabel = unprintable | keyLabel(keyKPColon)
	keylabelKPHash             keyLabel = unprintable | keyLabel(keyKPHash)
	keylabelKPSpace            keyLabel = unprintable | keyLabel(keyKPSpace)
	keylabelKPAt               keyLabel = unprintable | keyLabel(keyKPAt)
	keylabelKPExclam           keyLabel = unprintable | keyLabel(keyKPExclam)
	keylabelKPMemStore         keyLabel = unprintable | keyLabel(keyKPMemStore)
	keylabelKPMemRecall        keyLabel = unprintable | keyLabel(keyKPMemRecall)
	keylabelKPMemClear         keyLabel = unprintable | keyLabel(keyKPMemClear)
	keylabelKPMemAdd           keyLabel = unprintable | keyLabel(keyKPMemAdd)
	keylabelKPMemSubtract      keyLabel = unprintable | keyLabel(keyKPMemSubtract)
	keylabelKPMemMultiply      keyLabel = unprintable | keyLabel(keyKPMemMultiply)
	keylabelKPMemDivide        keyLabel = unprintable | keyLabel(keyKPMemDivide)
	keylabelKPPlusMinus        keyLabel = unprintable | keyLabel(keyKPPlusMinus)
	keylabelKPClear            keyLabel = unprintable | keyLabel(keyKPClear)
	keylabelKPClearEntry       keyLabel = unprintable | keyLabel(keyKPClearEntry)
	keylabelKPBinary           keyLabel = unprintable | keyLabel(keyKPBinary)
	keylabelKPOctal            keyLabel = unprintable | keyLabel(keyKPOctal)
	keylabelKPDecimal          keyLabel = unprintable | keyLabel(keyKPDecimal)
	keylabelKPHexadecimal      keyLabel = unprintable | keyLabel(keyKPHexadecimal)

	keylabelLCtrl  keyLabel = unprintable | keyLabel(keyLCtrl)
	keylabelLShift keyLabel = unprintable | keyLabel(keyLShift)
	keylabelLAlt   keyLabel = unprintable | keyLabel(keyLAlt)
	keylabelLGUI   keyLabel = unprintable | keyLabel(keyLGUI)
	keylabelRCtrl  keyLabel = unprintable | keyLabel(keyRCtrl)
	keylabelRShift keyLabel = unprintable | keyLabel(keyRShift)
	keylabelRAlt   keyLabel = unprintable | keyLabel(keyRAlt)
	keylabelRGUI   keyLabel = unprintable | keyLabel(keyRGUI)

	keylabelMode keyLabel = unprintable | keyLabel(keyMode)

	keylabelAudioNext   keyLabel = unprintable | keyLabel(keyAudioNext)
	keylabelAudioPrev   keyLabel = unprintable | keyLabel(keyAudioPrev)
	keylabelAudioStop   keyLabel = unprintable | keyLabel(keyAudioStop)
	keylabelAudioPlay   keyLabel = unprintable | keyLabel(keyAudioPlay)
	keylabelAudioMute   keyLabel = unprintable | keyLabel(keyAudioMute)
	keylabelMediaSelect keyLabel = unprintable | keyLabel(keyMediaSelect)
	keylabelWWW         keyLabel = unprintable | keyLabel(keyWWW)
	keylabelMail        keyLabel = unprintable | keyLabel(keyMail)
	keylabelCalculator  keyLabel = unprintable | keyLabel(keyCalculator)
	keylabelComputer    keyLabel = unprintable | keyLabel(keyComputer)
	keylabelACSearch    keyLabel = unprintable | keyLabel(keyACSearch)
	keylabelACHome      keyLabel = unprintable | keyLabel(keyACHome)
	keylabelACBack      keyLabel = unprintable | keyLabel(keyACBack)
	keylabelACForward   keyLabel = unprintable | keyLabel(keyACForward)
	keylabelACStop      keyLabel = unprintable | keyLabel(keyACStop)
	keylabelACRefresh   keyLabel = unprintable | keyLabel(keyACRefresh)
	keylabelACBookmarks keyLabel = unprintable | keyLabel(keyACBookmarks)

	keylabelBrightnessDown keyLabel = unprintable | keyLabel(keyBrightnessDown)
	keylabelBrightnessUp   keyLabel = unprintable | keyLabel(keyBrightnessUp)
	keylabelDisplaySwitch  keyLabel = unprintable | keyLabel(keyDisplaySwitch)
	keylabelKbdIllumToggle keyLabel = unprintable | keyLabel(keyKbdIllumToggle)
	keylabelKbdIllumDown   keyLabel = unprintable | keyLabel(keyKbdIllumDown)
	keylabelKbdIllumUp     keyLabel = unprintable | keyLabel(keyKbdIllumUp)
	keylabelEject          keyLabel = unprintable | keyLabel(keyEject)
	keylabelSleep          keyLabel = unprintable | keyLabel(keySleep)
)

const unprintable = 1 << 30

//------------------------------------------------------------------------------
