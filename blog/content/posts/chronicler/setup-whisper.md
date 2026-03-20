---
title: "Setup Whisper"
date: 2025-07-15T08:24:28+01:00
categories:
  - chronicler
tags:
  - chronicler
  - python
  - arch
  - pytorch
  - cuda
---

I wanted to generate chapter markers from a devlog audio recording using
OpenAI's Whisper, and figured I'd run it locally. Whisper is Python-based, and
I'm on Arch. What could go wrong?

Turns out… not much, but it still took a few hops.

## Choosing the Right Setup

I already had Python installed, but rather than littering system Python or
managing a bunch of ad hoc virtualenvs, I decided to do it properly — with
Poetry.

```bash
sudo pacman -S poetry
poetry new whisper-transcriber
cd whisper-transcriber
```

So far so good.

## PyTorch + CUDA: the PyPy Pitfall

My first attempt to install `torch`, `torchvision`, and `torchaudio` failed in a
confusing way — no versions found at all. The clue was in the command: I'd
accidentally run it with `pip-pypy3`. PyTorch doesn't build wheels for PyPy.
CPython only.

## Sorting Out Python Versions

My system Python was 3.13. PyTorch had just released 3.13 wheels for `torch`,
but not yet for `torchaudio` — version mismatch. I used `pyenv` to install 3.12
instead:

```bash
pyenv install 3.12.3
```

Updated `pyproject.toml`:

```toml
python = ">=3.12,<3.14"
```

And re-pointed Poetry:

```bash
poetry env use $(pyenv prefix 3.12.3)/bin/python
```

Poetry ignored me the first time because 3.13 was still hardcoded. After
recreating the environment and verifying the version, I was ready.

## PEP 668: the "Externally Managed" False Alarm

Even inside the Poetry shell, Arch's patched Python threw a
`--break-system-packages` error. This check is meant to protect system Python —
but it was firing inside a fully isolated Poetry environment. Safe to ignore. I
added the flag:

```bash
poetry run pip install torch torchvision torchaudio \
  --index-url https://download.pytorch.org/whl/cu121 \
  --break-system-packages
```

Worked perfectly.

## The Result

```bash
poetry run whisper output000.mp3 --model base --output_format json
```

Transcribed 2.5 hours of audio, timestamped segments ready for chapter
generation. All local, GPU-accelerated, isolated from system Python, and
repeatable.

---

## In Summary

If you're on Arch and want Whisper with CUDA:

1. Use `poetry` + `pyenv`
2. Set Python to 3.12 (not 3.13)
3. Install torch with `--break-system-packages` and the `cu121` index
4. Whisper just works

A bit of fiddling up front, but now it's a solid local tool — one less cloud
dependency to think about.
