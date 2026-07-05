# Video Merger

EN · A Windows desktop program (Go + Wails) that finds numbered video parts, checks whether they can be joined safely, previews the combined timeline, and merges them with FFmpeg stream copy.

KO · 번호가 붙은 동영상 조각을 자동으로 찾고, 안전하게 이어 붙일 수 있는지 확인한 뒤, 합쳐질 타임라인을 미리 보여 주고 FFmpeg 스트림 복사 방식으로 병합하는 Windows 데스크톱 프로그램입니다.

> Built with the help of Claude and ChatGPT (Vibe coding). This is a personal project, so bugs may still exist.

---

## English

### Screenshot

<img width="1226" height="1233" alt="English" src="https://github.com/user-attachments/assets/ca123086-7768-4203-8f22-da2c81f600ec" />

### What it is

Video Merger is a small Windows utility for joining split video files that already belong together.

It is designed for cases where one video has been divided into numbered parts, such as:

- `Lecture (1).mp4`, `Lecture (2).mp4`, `Lecture (3).mp4`
- `Camera_1.mov`, `Camera_2.mov`
- `Session - 1.mkv`, `Session - 2.mkv`

The program scans selected files or folders, groups matching parts, checks compatibility with FFprobe, shows the planned output, and then runs FFmpeg concat merging when the user starts the merge.

### Main features

- **Batch grouping**: detects numbered parts and groups them by folder, shared name, and extension.
- **Multiple numbering styles**: supports `name (1)`, `name [1]`, `name_1`, `name - 1`, `name 1`, and `name.1`.
- **Custom detection pattern**: advanced users can provide a regular expression with capture groups for shared name and part number.
- **Compatibility analysis**: checks stream count, codec, dimensions, pixel format, frame rate, audio sample rate, channel count, and channel layout metadata.
- **Preview**: lets the user play the planned combined timeline before producing the final output.
- **Stream-copy merge**: uses FFmpeg concat with `-c copy`, so compatible files can be joined without re-encoding.
- **Single-file handling**: a group containing only one file is copied instead of being merged.
- **Output planning**: shows the planned output name and prevents accidental overwrite by finding an available filename.
- **Progress reporting**: reports group progress, clip count, size, duration, loudness, and merge status.

### What it does not include

This repository does not bundle external media tools (FFmpeg).

FFmpeg and FFprobe must be available on the user’s machine, normally through `PATH`.

### Supported input extensions

The current file scanner accepts these video extensions:

- `.mp4`
- `.mov`
- `.mkv`
- `.m4v`

Other formats may still be supported by FFmpeg itself, but they are not currently included in the program’s scan filter.

### How it works

1. Select source files or source folders.
2. Choose an output folder, or choose the same folder as the input.
3. Choose whether to include subfolders.
4. Run **Analyze**.
5. Review detected groups, notices, cautions, warnings, planned output, and preview timeline.
6. Enable **Ignore cautions** only when the cautions are acceptable.
7. Enable **Force merge** only when you intentionally want FFmpeg to try merging despite warnings.
8. Run **Merge**.
9. Check the output files and final result report.

### Merge model

Video Merger uses FFmpeg concat demuxer merging with stream copy:

```powershell
ffmpeg -f concat -safe 0 -i <concat-list> -c copy <output-file>
```

This is fast and preserves the existing video/audio streams, but it also means the input parts should be structurally compatible. When the streams differ, FFmpeg may fail or produce a bad result. That is why the program emphasizes analysis before merging.

### Requirements

For normal use:

- Windows 10 or later.
- FFmpeg available from the command line as `ffmpeg`.
- FFprobe available from the command line as `ffprobe`.
- Microsoft Edge WebView2 Runtime, normally already present on modern Windows systems.

For development:

- Go 1.23 or later.
- Wails CLI v2.
- PowerShell.

### Build from source

Install Wails first if it is not already available:

```powershell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### License

This project is released under the MIT License. See [`LICENSE`](./LICENSE) for details.

---

## 한국어

### 스크린샷

<img width="1226" height="1233" alt="Korean" src="https://github.com/user-attachments/assets/dceca962-c290-4d4b-a8b9-0216dc4f5c32" />

### 이 프로그램은?

Video Merger는 분할 동영상 파일을 합치는 작은 Windows 유틸리티입니다.

다음처럼 번호가 붙어 있는 영상을 자동으로 인식하여 하나로 합칩니다.

- `Lecture (1).mp4`, `Lecture (2).mp4`, `Lecture (3).mp4`
- `Camera_1.mov`, `Camera_2.mov`
- `Session - 1.mkv`, `Session - 2.mkv`

선택한 폴더를 스캔하여 하나로 합쳐야 할 영상을 인식하고, FFprobe로 호환성을 확인하여, FFmpeg concat 병합을 수행합니다.

### 주요 기능

- **그룹화**: 번호가 붙은 파일을 감지하고 폴더, 공통된 이름, 확장자를 기준으로 그룹화합니다.
- **여러 번호 형식**: `name (1)`, `name [1]`, `name_1`, `name - 1`, `name 1`, `name.1` 형식을 지원합니다.
- **사용자 지정 감지 패턴**: 공통된 이름과 번호를 인식하는 정규식을 지정할 수도 있습니다.
- **호환성 분석**: 스트림 수, 코덱, 해상도, 가로 세로 크기, 프레임 레이트, 오디오 샘플 레이트, 채널 수, 채널 레이아웃 메타데이터를 확인하여 안전하게 합칠 수 있는지 확인합니다.
- **미리보기**: 합쳤을 경우 타임라인을 실제로 합치지 않고 보여주는 미리보기를 제공합니다.
- **스트림 복사 병합**: FFmpeg concat과 `-c copy`를 사용하여 호환되는 파일은 재인코딩 없이 이어 붙입니다.
- **단일 파일 처리**: 파일이 하나뿐인 그룹은 병합하지 않고 복사합니다.
- **출력 계획**: 최종 파일 이름을 보여 주어 실수로 덮어쓰는 일을 예방합니다.
- **진행 상황 표시**: 진행률, 클립 수, 크기, 길이, 음량, 병합 상태를 보여줍니다.

### 다음은 포함되어 있지 않습니다

이 프로그램은 FFmpeg 실행 파일이 포함되어 있지 않습니다.

FFmpeg는 사전에 설치되어 있어야 합니다 (`PATH`를 통해 찾을 수 있어야 합니다.)

### 지원 입력 확장자

현재 다음의 확장자를 지원합니다.

- `.mp4`
- `.mov`
- `.mkv`
- `.m4v`

FFmpeg에서 지원하는 다른 형식의 아직 지원하지 않습니다.

### 작동 방식

1. 원본 파일이나 원본 폴더를 선택합니다.
2. 저장 폴더를 선택하거나 원본과 같은 폴더를 선택합니다.
3. 하위 폴더를 포함할지 선택합니다.
4. **분석**을 실행합니다.
5. 합칠 파일들을 확인합니다.
6. 주의 사항을 무시하고자 할 경우 **주의 무시**를 선택할 수 있습니다.
7. 경고를 무시하고 합쳐야 할 경우에만 **강제 병합**을 선택해 주세요.
8. **병합**을 실행합니다.
9. 출력 파일과 최종 결과 보고를 확인합니다.

### 병합 방식

Video Merger는 FFmpeg concat demuxer와 스트림 복사를 사용해 영상 파일을 합칩니다.

```powershell
ffmpeg -f concat -safe 0 -i <concat-list> -c copy <output-file>
```

기존 비디오/오디오 스트림을 사용하여 빠르게 합칠 수 있으나, 합칠 파일들이 호환되어야 합니다. 호환되지 않을 경우 실패하거나 잘못된 파일을 만들 수 있습니다.

### 요구 사항

일반 사용:

- Windows 10 이상.
- 명령줄에서 `ffmpeg`로 사용할 수 있는 FFmpeg.
- 명령줄에서 `ffprobe`로 사용할 수 있는 FFprobe.
- Microsoft Edge WebView2 Runtime. 최신 Windows 시스템에는 보통 이미 설치되어 있습니다.

개발:

- Go 1.23 이상.
- Wails CLI v2.
- PowerShell.

### 소스에서 빌드

Wails가 아직 설치되어 있지 않다면 먼저 설치합니다.

```powershell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 라이선스

이 프로젝트는 MIT License로 배포됩니다. 자세한 내용은 [`LICENSE`](./LICENSE)를 확인하세요.
