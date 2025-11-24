# Notify

[English](#english) | [ภาษาไทย](#thai)

---

<a name="english"></a>
## English

**Notify** is a professional, secure, and easy-to-use Go library for sending notifications to various messaging platforms. It is designed to be stateless (no hardcoded secrets) and supports both **simple cross-platform messages** and **advanced provider-specific features** (Flex Messages, Adaptive Cards, Embeds).

### Features
- **Multi-Provider Support**: LINE (Messaging API), Telegram, Discord, Microsoft Teams.
- **Flexible Interface**: Send simple text, generic rich messages, or full API payloads.
- **Advanced Features**:
    - **LINE**: Flex Messages, Templates, Quick Replies.
    - **Telegram**: Keyboards, ParseMode (Markdown/HTML).
    - **Discord**: Rich Embeds (Fields, Footer, Author), Webhook customization.
    - **MS Teams**: Full Adaptive Cards support.
- **Professional**: Functional Options pattern, Context support, Unit Tested.

### Installation

```bash
go get github.com/thanpawatpiti/notify
```

### Usage

#### 1. Import the library
```go
import (
    "context"
    "github.com/thanpawatpiti/notify"
    "github.com/thanpawatpiti/notify/providers/line"
    // ... other providers
)
```

#### 2. Send Notifications

**Simple Text (All Providers)**
```go
p.Send(ctx, "Hello World")
```

**Common Rich Message (All Providers)**
```go
msg := notify.CommonMessage{
    Title:    "Hello",
    Content:  "Rich content",
    ImageURL: "https://example.com/image.png",
}
p.Send(ctx, msg)
```

**Advanced: LINE Flex Message**
```go
flexMsg := line.FlexMessage{
    AltText: "Flex Message",
    Contents: line.BubbleContainer{
        Type: "bubble",
        Body: &line.BoxComponent{
            Type: "box",
            Layout: "vertical",
            Contents: []line.FlexComponent{
                line.TextComponent{Type: "text", Text: "Hello Flex!"},
            },
        },
    },
}
lineProvider.Send(ctx, flexMsg)
```

**Advanced: Discord Embed**
```go
embed := discord.Embed{
    Title: "Advanced Embed",
    Fields: []discord.EmbedField{
        {Name: "Field 1", Value: "Value 1", Inline: true},
    },
}
discordProvider.Send(ctx, embed)
```

---

<a name="thai"></a>
## ภาษาไทย

**Notify** คือไลบรารีภาษา Go สำหรับส่งการแจ้งเตือนที่ยืดหยุ่นและเป็นมืออาชีพ รองรับทั้งการส่งข้อความแบบง่ายๆ และการใช้ฟีเจอร์ขั้นสูงของแต่ละแพลตฟอร์ม

### คุณสมบัติ
- **รองรับหลายแพลตฟอร์ม**: LINE, Telegram, Discord, MS Teams
- **ยืดหยุ่น**: ส่งข้อความได้ทั้งแบบ Text ธรรมดา, แบบ Rich Message ทั่วไป, หรือแบบโครงสร้างเฉพาะของแต่ละค่าย (Payload)
- **ฟีเจอร์ขั้นสูง**:
    - **LINE**: รองรับ Flex Message เต็มรูปแบบ
    - **Telegram**: รองรับ Keyboard, Markdown/HTML
    - **Discord**: รองรับ Embed แบบเต็มสูบ (Fields, Footer)
    - **MS Teams**: รองรับ Adaptive Cards
- **มืออาชีพ**: ออกแบบด้วย Functional Options Pattern

### วิธีใช้งาน

#### 1. ส่งข้อความธรรมดา (ทุกค่าย)
```go
p.Send(ctx, "สวัสดีครับ")
```

#### 2. ส่งข้อความแบบ Rich Message (ทุกค่าย)
```go
msg := notify.CommonMessage{
    Title:    "สวัสดี",
    Content:  "ข้อความพร้อมรูปภาพ",
    ImageURL: "https://example.com/image.png",
}
p.Send(ctx, msg)
```

#### 3. ขั้นสูง: LINE Flex Message
```go
flexMsg := line.FlexMessage{
    AltText: "ตัวอย่าง Flex",
    Contents: line.BubbleContainer{
        Type: "bubble",
        Body: &line.BoxComponent{
            Type: "box",
            Layout: "vertical",
            Contents: []line.FlexComponent{
                line.TextComponent{Type: "text", Text: "สวัสดี Flex!"},
            },
        },
    },
}
lineProvider.Send(ctx, flexMsg)
```
