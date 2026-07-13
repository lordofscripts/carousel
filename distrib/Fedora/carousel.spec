%global debug_package %{nil} %global _builddebuginfo_packages %{nil}
%define bindir /usr/local/bin
%define __cp %{__cp} -r

Name:           carousel
Version:        1.1.0-RC.1
Release:        1%{?dist}
Summary:        Carousel is a versatile multi-platform Wallpaper sequence.

License:        MIT
URL:            https://github.com/lordofscripts/carousel
Source:         %{name}-%{version}.tar.gz
BuildArch:      x86_64
BuildRequires:  golang
Requires:       golang

%description
Carousel is a versatile multi-platform Wallpaper sequencer with several features.

%prep
%setup -q

%build
[ -d bin ] || mkdir bin
go build -tags logx -v -buildmode=pie -o bin/goCarousel cmd/*.go

%install
rm -rf $RPM_BUILD_ROOT
mkdir -p $RPM_BUILD_ROOT/%{_bindir}
install -m 0755  bin/goCarousel $RPM_BUILD_ROOT/%{_bindir}
#mkdir -p $RPM_BUILD_ROOT/%{_sysconfdir}
#install %{name}rc $RPM_BUILD_ROOT/%{_sysconfdir}

%postun
rm -f $RPM_BUILD_ROOT/%{_bindir}/goCarousel

%clean
rm -rf $RPM_BUILD_ROOT

%files
%{_bindir}/goCarousel
#%{_sysconfdir}/%{name}rc
%license LICENSE.md

%changelog
* Mon Jul 14 2026 lordofscripts
- First RPM package (v1.1.0-RC.1)

* Sat Sep 27 2025 lordofscripts
- initial release
