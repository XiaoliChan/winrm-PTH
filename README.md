<a name="readme-top"></a>

# winrm-PTH
Golang implement WinRM client with PTH (pass the hash)

#### Table of Contents
<ol>
<li>
  <a href="#info">Info</a>
</li>
<li><a href="#how-it-works">How it works?</a></li>
<li><a href="#screenshots">Screenshots</a></li>
<li><a href="#references">References</a></li>
</ol>

## Info

This is an example of Golang implement WinRM client with PTH.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How it works?

In [ntlmssp](https://github.com/bodgit/ntlmssp), it always encodes the password to md4 hash in function `ntowfV1`, then we can make a simple logic detection to check if the user passes a hash into it, with that "patch", we can give hash into `ntlmssp` to make the auth structures.

The [winrm](https://github.com/masterzen/winrm) always use `DefaultClientCompatibilityLevel`, so we don't need to do lots of changes in `ntlmssp`.

Why do we need `encryption.go`?

- Because in Windows, winrm is always set `AllowUnencrypted` to false, and the library [winrm](https://github.com/masterzen/winrm) only support `AllowUnencrypted`, which means we can't auth into `winrm` without `winrm set winrm/config/service @{AllowUnencrypted="true"}`, this is not make sense.
- With `encryption.go`, we can auth into `winrm` even without `winrm set winrm/config/service @{AllowUnencrypted="true"}`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Screenshots
- Without `AllowUnencrypted` in PTH
![image](https://github.com/XiaoliChan/winrm-PTH/assets/30458572/12f94519-6451-471d-9880-80eb4cd4fc28)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## References
- https://github.com/CalypsoSys/winrmntlm
- https://github.com/masterzen/winrm
- https://github.com/CalypsoSys/bobwinrm
- https://github.com/bodgit/ntlmssp

<p align="right">(<a href="#readme-top">back to top</a>)</p>
